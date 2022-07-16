package datarepository

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

type fileRepo struct {
}

// NewFileRepository create an handle to a File reader and writer
func NewDBRepository() {
	return
}
