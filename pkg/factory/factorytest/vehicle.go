package factorytest

import "github.com/rolandjitsu/go-design-patterns/pkg/factory"

type mockVehicle struct {
	b  string
	m  string
	sn int
}

func (m *mockVehicle) Brand() string {
	return m.b
}

func (m *mockVehicle) Model() string {
	return m.m
}

func (m *mockVehicle) SN() int {
	return m.sn
}

var _ factory.Vehicle = &mockVehicle{}
