package datastore

//Port an object to transport Port data to and from a data repository
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
