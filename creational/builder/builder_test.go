package builder

import "testing"

func TestContructMethod(t *testing.T) {
	manufactureDirector := new(ManufacturingDirector)
	carBuilder := new(CarBuilder)

	manufactureDirector.SetBuilder(carBuilder)
	manufactureDirector.Construct()
	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheel on car must be 4 and there are %d wheels\n", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Errorf("Structure of car must be 'Car' and was %s\n", car.Structure)
	}

	if car.Seats != 4 {
		t.Errorf("Seat on car must be 4 and there are %d seats\n", car.Seats)
	}

	vanBuilder := new(VanBuilder)
	manufactureDirector.SetBuilder(vanBuilder)
	manufactureDirector.Construct()
	van := vanBuilder.GetVehicle()

	if van.Wheels != 4 {
		t.Errorf("Wheel on van must be 4 and there are %d wheels\n", van.Wheels)
	}

	if van.Structure != "Van" {
		t.Errorf("Structure of van must be 'Van' and was %s\n", van.Structure)
	}

	if van.Seats != 8 {
		t.Errorf("Seat on van must be 8 and there are %d seats\n", van.Seats)
	}	

}
