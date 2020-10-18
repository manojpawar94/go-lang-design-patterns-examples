package builder

//CarBuilder struct
type CarBuilder struct {
	Vehicle Vehicle
}

//SetWheels method
func (c *CarBuilder) SetWheels() BuildProcess {
	c.Vehicle.Wheels = 4
	return c
}

//SetSeats method
func (c *CarBuilder) SetSeats() BuildProcess {
	c.Vehicle.Seats = 4
	return c
}

//SetStructure method
func (c *CarBuilder) SetStructure() BuildProcess {
	c.Vehicle.Structure = "Car"
	return c
}

//GetVehicle method
func (c *CarBuilder) GetVehicle() Vehicle {
	return c.Vehicle
}