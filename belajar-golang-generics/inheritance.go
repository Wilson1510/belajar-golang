package main

import (
	"fmt"
)

type Employee interface {
	GetName() string
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

func GetName[T Employee](employee T) string {
	return employee.GetName()
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresidentName() string {
	return m.Name
}

func main() {
	manager := GetName[Manager](&MyManager{Name: "John"})
	fmt.Println(manager)

	vicePresident := GetName[VicePresident](&MyVicePresident{Name: "Jane"})
	fmt.Println(vicePresident)
}
