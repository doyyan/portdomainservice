package interactor

import (
	"github.com/doyyan/portdomainservice/domain/model"
	"github.com/doyyan/portdomainservice/usecase/presenter"
	"github.com/doyyan/portdomainservice/usecase/repository"
)

type portInteractor struct {
	DBRepository  repository.DBRepository
	PortPresentor presenter.PortsPresenter
}

type PortInterActor interface {
	// CreatePorts creates Ports in the datastore
	CreatePorts(ports []*model.Port) ([]*model.Port, error)
	// Updates Ports in the datastore
	UpdatePorts(ports []*model.Port) ([]*model.Port, error)
}

func NewPortInteractor(db repository.DBRepository, p presenter.PortsPresenter) portInteractor {
	return portInteractor{DBRepository: db, PortPresentor: p}
}

//CreatePorts calls the DB's create port
func (p portInteractor) CreatePorts(ports []*model.Port) ([]*model.Port, error) {
	var createdPorts []*model.Port
	for _, port := range ports {
		valid, err := port.Validate()
		if err != nil {
			return nil, err
		}
		doesExist, err := p.DBRepository.FindPort(*port)
		if err != nil {
			return nil, err
		}
		if valid && !doesExist {
			p.DBRepository.CreatePort(*port)
		}
		createdPorts = append(createdPorts, port)
	}
	return createdPorts, nil
}

//UpdatePorts updates a port
func (p portInteractor) UpdatePorts(ports []*model.Port) ([]*model.Port, error) {
	var updatedPorts []*model.Port
	for _, port := range ports {
		valid, err := port.Validate()
		if err != nil {
			return nil, err
		}
		doesExist, err := p.DBRepository.FindPort(*port)
		if err != nil {
			return nil, err
		}
		if valid && !doesExist {
			p.DBRepository.UpdatePort(*port)
		}
		updatedPorts = append(updatedPorts, port)
	}
	return updatedPorts, nil
}
