package presenter

import "github.com/doyyan/portdomainservice/domain/model"

//go:generate moq -out mocks/portpresenter.mock.go -pkg portpresenter_test -skip-ensure . PortsPresenter
type PortsPresenter interface {
	ResponsePorts(ports []*model.Port) []*model.Port
}
