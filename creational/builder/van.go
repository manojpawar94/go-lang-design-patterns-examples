package builder

//VanBuilder struct
type VanBuilder struct {
	Vehicle Vehicle
}

//SetWheels method
func (c *VanBuilder) SetWheels() BuildProcess {
	c.Vehicle.Wheels = 4
	return c
}

//SetSeats method
func (c *VanBuilder) SetSeats() BuildProcess {
	c.Vehicle.Seats = 8
	return c
}

//SetStructure method
func (c *VanBuilder) SetStructure() BuildProcess {
	c.Vehicle.Structure = "Van"
	return c
}

//GetVehicle method
func (c *VanBuilder) GetVehicle() Vehicle {
	return c.Vehicle
}