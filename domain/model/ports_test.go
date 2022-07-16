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
	tests := []struct {
		name string
		args args
		want *Port
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPort(tt.args.name, tt.args.city, tt.args.country, tt.args.alias, tt.args.regions, tt.args.coordinates, tt.args.province, tt.args.timezone, tt.args.unlocode, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
