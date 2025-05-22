package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("Not Implemented")
	ErrTruckNotFound  = errors.New("Truck Not Found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
	GetID() string
	GetCargo() int
}

type NormalTruck struct {
	id    string
	cargo int
}

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}
func (t *NormalTruck) UnloadCargo() error {
	t.cargo -= 2
	return nil
}

func (t *NormalTruck) GetID() string {
	return t.id
}
func (t *NormalTruck) GetCargo() int {
	return t.cargo
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (t ElectricTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}
func (t ElectricTruck) UnloadCargo() error {
	t.cargo -= 2
	return nil
}

func processTruck(truck Truck) error {
	fmt.Println("Processing truck with ID:", truck.GetID())

	return ErrNotImplemented
}

func main() {
	trucks := []NormalTruck{
		{id: "T1"},
		{id: "T2"},
		{id: "T3"},
	}

	for _, truck := range trucks {
		fmt.Printf("Truck %s is arrived.\n", truck.id)
		err := processTruck(&truck)
		if err != nil {
			if err == ErrNotImplemented {
				fmt.Printf("Truck %s is not implemented.\n", truck.id)
			}
			if errors.Is(err, ErrTruckNotFound) {
				fmt.Printf("Truck %s is not found.\n", truck.id)
			}

			log.Fatal("Error processing truck: %s", err)
		}
	}

}
