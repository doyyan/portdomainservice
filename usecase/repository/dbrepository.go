package repository

import "github.com/doyyan/portdomainservice/domain/model"

//go:generate moq -out mocks/derpeository.mock.go -pkg dbrepository_test -skip-ensure . DBRepository

type DBRepository interface {
	// CreatePorts creates Ports in the datastore
	CreatePort(port model.Port) (model.Port, error)
	// Updates Ports in the datastore
	UpdatePort(port model.Port) (model.Port, error)
	// Updates Ports in the datastore
	FindPort(port model.Port) (bool, error)
}
