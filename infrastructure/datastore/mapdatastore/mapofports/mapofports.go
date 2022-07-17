package mapofports

import "github.com/doyyan/portdomainservice/dto/datastore"

//MapOfPorts a data storage map for a list of Ports
var MapOfPorts map[string]datastore.Port = make(map[string]datastore.Port)

//NewMapOfPorts create a new instance of a data storage map for Ports
func NewMapOfPorts() map[string]datastore.Port {
	return MapOfPorts
}
