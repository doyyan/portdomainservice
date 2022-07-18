package db

import (
	"github.com/doyyan/portdomainservice/infrastructure/datastore/mapdatastore/mapofports"
	"github.com/doyyan/portdomainservice/interfaces/datarepository"
)

func NewDB() datarepository.DBRepo {
	database := mapofports.NewMapOfPorts()
	return datarepository.NewDBRepository(database)
}
