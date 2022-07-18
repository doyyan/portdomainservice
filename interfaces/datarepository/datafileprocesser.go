package datarepository

import (
	"github.com/doyyan/portdomainservice/domain/model"
	"github.com/doyyan/portdomainservice/dto/datastore"
)

//Port a port entity
type Port struct {
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocode    string
	Code        string
}

type DBRepo struct {
	Datastore map[string]datastore.Port
}

// NewFileRepository create an handle to a File reader and writer
func NewDBRepository(ports map[string]datastore.Port) DBRepo {
	return DBRepo{ports}
}

// CreatePort map a Port domain to a Port DB Entity and Create in DB, if port not found
func (d DBRepo) CreatePort(port model.Port) (model.Port, error) {
	newPort := datastore.Port{
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		//Unlocode:    port.Unlocode,
		Code: port.Code,
	}
	d.Datastore[port.Name] = newPort
	return port, nil
}

// UpdatePort map a Port domain to a Port DB Entity and Update in DB, if port found
func (d DBRepo) UpdatePort(port model.Port) (model.Port, error) {
	found, _ := d.FindPort(port)
	if found {
		updatedPort := datastore.Port{
			Name:        port.Name,
			City:        port.City,
			Country:     port.Country,
			Alias:       port.Alias,
			Regions:     port.Regions,
			Coordinates: port.Coordinates,
			Province:    port.Province,
			Timezone:    port.Timezone,
			//		Unlocode:    port.Unlocode,
			Code: port.Code,
		}
		d.Datastore[port.Name] = updatedPort
		return port, nil
	}
	return model.Port{}, nil
}

// FindPort find if the Port exists in the DB
func (d DBRepo) FindPort(port model.Port) (bool, error) {
	_, ok := d.Datastore[port.Name]
	return ok, nil
}
