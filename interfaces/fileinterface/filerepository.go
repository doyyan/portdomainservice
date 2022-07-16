package fileinterface

import "github.com/doyyan/portdomainservice/dto/datastore"

type (
	//FileRepository an interface to talk to the DB/datastore
	FileRepository interface {
		// CreatePorts save a given list of Ports to the datastore
		CreatePorts(ports []*datastore.Port) ([]*datastore.Port, error)
		// UpdatePorts update data for a given list of ports in the datastore
		UpdatePorts(ports []*datastore.Port) ([]*datastore.Port, error)
		// FindAPort find a single Port in the datastore
		FindAPort(ports datastore.Port) (datastore.Port, error)
	}
)
