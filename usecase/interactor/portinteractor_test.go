package interactor

import (
	"reflect"
	"testing"

	"github.com/doyyan/portdomainservice/domain/model"
	"github.com/doyyan/portdomainservice/usecase/presenter"
	portpresenter_test "github.com/doyyan/portdomainservice/usecase/presenter/mocks"
	"github.com/doyyan/portdomainservice/usecase/repository"
	dbrepository_test "github.com/doyyan/portdomainservice/usecase/repository/mocks"
)

func TestNewPortInteractor(t *testing.T) {
	type args struct {
		db repository.DBRepository
		p  presenter.PortsPresenter
	}
	tests := []struct {
		name string
		args args
		want *portInteractor
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPortInteractor(tt.args.db, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPortInteractor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_portInteractor_CreatePorts(t *testing.T) {
	type fields struct {
		DBRepository  repository.DBRepository
		PortPresentor presenter.PortsPresenter
	}
	type args struct {
		ports []*model.Port
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Port
		wantErr bool
	}{
		{
			name: "Success Update Port",
			fields: fields{
				DBRepository:  dbRepo(),
				PortPresentor: portPresenter(),
			},
			args: struct{ ports []*model.Port }{ports: testPortPointers(testPorts())},
			want: testPortPointers(testPorts()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &portInteractor{
				DBRepository:  tt.fields.DBRepository,
				PortPresentor: tt.fields.PortPresentor,
			}
			got, err := p.CreatePorts(tt.args.ports)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePorts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePorts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_portInteractor_UpdatePorts(t *testing.T) {
	type fields struct {
		DBRepository  repository.DBRepository
		PortPresentor presenter.PortsPresenter
	}
	type args struct {
		ports []*model.Port
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.Port
		wantErr bool
	}{
		{
			name: "Success Update Port",
			fields: fields{
				DBRepository:  dbRepo(),
				PortPresentor: portPresenter(),
			},
			args: struct{ ports []*model.Port }{ports: testPortPointers(testPorts())},
			want: testPortPointers(testPorts()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &portInteractor{
				DBRepository:  tt.fields.DBRepository,
				PortPresentor: tt.fields.PortPresentor,
			}
			got, err := p.UpdatePorts(tt.args.ports)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdatePorts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdatePorts() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func testPorts() []model.Port {
	return []model.Port{model.Port{
		Name:        "Port1",
		City:        "City1",
		Country:     "Country1",
		Alias:       nil,
		Regions:     nil,
		Coordinates: []float64{1.2, 3.4},
		Province:    "",
		Timezone:    "",
		Unlocode:    "",
		Code:        "",
	}}
}

func testPortPointers(port []model.Port) []*model.Port {
	portAddresses := []*model.Port{}
	for _, p := range port {
		portAddresses = append(portAddresses, &p)
	}
	return portAddresses
}

func portPresenter() portpresenter_test.PortsPresenterMock {
	port := testPortPointers(testPorts())
	p := portpresenter_test.PortsPresenterMock{ResponsePortsFunc: func(ports []*model.Port) []*model.Port {
		return port
	}}
	return p
}
func dbRepo() dbrepository_test.DBRepositoryMock {
	d := dbrepository_test.DBRepositoryMock{
		CreatePortFunc: func(port model.Port) (model.Port, error) {
			return testPorts()[0], nil
		},
		FindPortFunc: func(port model.Port) (bool, error) {
			return true, nil
		},
		UpdatePortFunc: func(port model.Port) (model.Port, error) {
			return testPorts()[0], nil
		},
	}
	return d
}
