package builder

//Vehicle struct
type Vehicle struct {
	Wheels    int
	Seats     int
	Structure string
}

//BuildProcess interface
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() Vehicle
}

//ManufacturingDirector struct
type ManufacturingDirector struct {
	builder BuildProcess
}

//SetBuilder method
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

//Construct method
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}
