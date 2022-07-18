package portcontroller

import (
	"github.com/doyyan/portdomainservice/interfaces/controller"
	"github.com/doyyan/portdomainservice/interfaces/datarepository"
	"github.com/doyyan/portdomainservice/interfaces/fileinterface"
	"github.com/doyyan/portdomainservice/usecase/interactor"
)

func NewPortController(db datarepository.DBRepo) controller.PortController {
	p := interactor.NewPortInteractor(db, nil)
	c := controller.NewPortController(p, fileinterface.NewFilePortRepository())
	return c
}
