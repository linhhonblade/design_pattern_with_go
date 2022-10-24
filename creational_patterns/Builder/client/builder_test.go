package main

import (
	builder "github.com/linhhonblade/builder"
	"testing"
)

func TestBuilder(t *testing.T) {
	// Setup
	director := builder.ManufacturingDirector{}
	carBuilder := &builder.CarBuilder{}
	bikeBuilder := &builder.BikeBuilder{}

	// Build a car
	director.SetBuilder(carBuilder)
	car := director.Construct()
	if car.Wheels != 4 {
		t.Errorf("Wheels on car is %d, expected 4\n", car.Wheels)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on car is %d, expected 5\n", car.Seats)
	}
	if car.Structure != "Car" {
		t.Errorf("Car structure is %s, expected \"Car\"", car.Structure)
	}

	// Build a bike
	director.SetBuilder(bikeBuilder)
	bike := director.Construct()
	if bike.Wheels != 2 {
		t.Errorf("Wheels on bike is %d, expected 2\n", bike.Wheels)
	}
	if car.Seats != 5 {
		t.Errorf("Seats on bike is %d, expected 2\n", bike.Seats)
	}
	if car.Structure != "Bike" {
		t.Errorf("Bike structure is %s, expected \"Bake\"", bike.Structure)
	}
}
