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
	CreatePort(ports *model.Port) ([]*model.Port, error)
	// Updates Ports in the datastore
	UpdatePort(ports *model.Port) ([]*model.Port, error)
}

func NewPortInteractor(db repository.DBRepository, p presenter.PortsPresenter) *portInteractor {
	return &portInteractor{DBRepository: db, PortPresentor: p}
}

//CreatePorts calls the DB's create port
func (p *portInteractor) CreatePorts(ports []*model.Port) ([]*model.Port, error) {
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
	}
	return ports, nil
}

//UpdatePorts updates a port
func (p *portInteractor) UpdatePorts(ports []*model.Port) ([]*model.Port, error) {
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
	}
	return ports, nil
}
