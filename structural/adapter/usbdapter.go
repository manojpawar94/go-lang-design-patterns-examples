package adapter

//Computer interface
type Computer interface {
	InsertInSqurePort()
}

//MacBook struct
type MacBook struct{}

//InsertInSqurePort method
func (*MacBook) InsertInSqurePort() {
	println("Inserting in square port in Macbook machine.")
}

//WindowPortAdpater stuct
type WindowPortAdpater struct {
	*Windows
}

//InsertInSqurePort method
func (w *WindowPortAdpater) InsertInSqurePort() {
	print("Converting Square port into Circle port... ")
	w.Windows.InsertInCirclePort()
}

//Windows struct
type Windows struct{}

//InsertInCirclePort method
func (*Windows) InsertInCirclePort() {
	println("Inserting in Circle port into Windows machine")
}

//Clinet struct
type Clinet struct{}

//InsetSquareUsbInComputer function
func (*Clinet) InsetSquareUsbInComputer(computer Computer) {
	computer.InsertInSqurePort()
}
