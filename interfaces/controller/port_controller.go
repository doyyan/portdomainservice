package controller

import (
	"github.com/doyyan/portdomainservice/domain/model"
	"github.com/doyyan/portdomainservice/dto/datastore"
	"github.com/doyyan/portdomainservice/usecase/interactor"
)

type PortController interface {
	CreatePorts() error
	UpdatePorts() error
}

type PortsDataProvider interface {
	// CreatePorts save a given list of Ports to the datastore
	CreatePorts() ([]*datastore.Port, error)
	// UpdatePorts update data for a given list of ports in the datastore
	UpdatePorts() ([]*datastore.Port, error)
}

type portController struct {
	portInteractor interactor.PortInterActor
	dataProvider   PortsDataProvider
}

func (p portController) CreatePorts() error {
	data, _ := p.dataProvider.CreatePorts()
	_, err := p.portInteractor.CreatePorts(dtoPortToDomainPort(data))
	return err
}

func (p portController) UpdatePorts() error {
	data, _ := p.dataProvider.UpdatePorts()
	_, err := p.portInteractor.UpdatePorts(dtoPortToDomainPort(data))
	return err
}

func NewPortController(portInteractor interactor.PortInterActor, dataProvider PortsDataProvider) PortController {
	return &portController{portInteractor: portInteractor, dataProvider: dataProvider}
}

func dtoPortToDomainPort(ports []*datastore.Port) []*model.Port {
	var allports []*model.Port
	for _, port := range ports {
		p := model.Port{
			Name:        port.Name,
			City:        port.City,
			Country:     port.Country,
			Alias:       []string(port.Alias),
			Regions:     port.Regions,
			Coordinates: port.Coordinates,
			Province:    port.Province,
			Timezone:    port.Timezone,
			Code:        port.Code,
		}
		allports = append(allports, &p)
	}
	return allports
}
