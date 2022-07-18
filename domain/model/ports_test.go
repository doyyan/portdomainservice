package model

import (
	"reflect"
	"testing"
)

func TestNewPort(t *testing.T) {
	type args struct {
		name        string
		city        string
		country     string
		alias       []string
		regions     []string
		coordinates []float64
		province    string
		timezone    string
		unlocode    string
		code        string
	}
	tests := map[string]struct {
		name string
		port Port
		want *Port
	}{
		"pass GoodPort": {
			port: Port{
				Name:        "Port1",
				City:        "City1",
				Country:     "Country1",
				Coordinates: []float64{1.2, 3.4},
			},
			want: &Port{
				Name:        "Port1",
				City:        "City1",
				Country:     "Country1",
				Coordinates: []float64{1.2, 3.4},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := NewPort(tt.port.Name, tt.port.City, tt.port.Country, tt.port.Alias, tt.port.Regions, tt.port.Coordinates,
				tt.port.Province, tt.port.Timezone, tt.port.Unlocode, tt.port.Code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
