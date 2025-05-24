package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	// Test cases for the NormalTruck
	normalTruck := &NormalTruck{id: "NT123", cargo: 0}
	if err := normalTruck.LoadCargo(); err != nil {
		t.Errorf("Failed to load cargo: %v", err)
	}
	if normalTruck.GetCargo() != 2 {
		t.Errorf("Expected cargo to be 2, got %d", normalTruck.GetCargo())
	}
	if err := normalTruck.UnloadCargo(); err != nil {
		t.Errorf("Failed to unload cargo: %v", err)
	}
	if normalTruck.GetCargo() != 0 {
		t.Errorf("Expected cargo to be 0, got %d", normalTruck.GetCargo())
	}

	// Test cases for the ElectricTruck
	electricTruck := &ElectricTruck{id: "ET456", cargo: 0, battery: 100.0}
	if err := electricTruck.LoadCargo(); err != nil {
		t.Errorf("Failed to load cargo: %v", err)
	}

}

// Test cases for the processTruck function
