package fileinterface

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/doyyan/portdomainservice/dto/datastore"
)

type (
	//FileRepository an interface to talk to the DB/datastore
	FileRepository interface {
		// CreatePorts save a given list of Ports to the datastore
		CreatePorts() ([]*datastore.Port, error)
		// UpdatePorts update data for a given list of ports in the datastore
		UpdatePorts() ([]*datastore.Port, error)
	}
)

type filedata struct{}

func NewFilePortRepository() filedata {
	return filedata{}
}
func (f filedata) CreatePorts() ([]*datastore.Port, error) {
	return readFile(), nil
}

func (f filedata) UpdatePorts() ([]*datastore.Port, error) {
	return readFile(), nil
}

type AllPorts struct {
	Ports []datastore.Port
}

func readFile() []*datastore.Port {
	var ports AllPorts
	// Open our jsonFile
	jsonFile, err := os.Open("ports.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &ports)
	var listOfPorts []*datastore.Port
	// append operation asynchronised to make it quicker
	var wg sync.WaitGroup
	wg.Add(len(ports.Ports))
	for _, port := range ports.Ports {
		go func() {
			listOfPorts = append(listOfPorts, &port)
			wg.Done()
		}()
	}
	wg.Wait()
	return listOfPorts
}
