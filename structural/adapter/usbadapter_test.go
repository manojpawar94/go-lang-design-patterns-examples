package adapter

import "testing"

func TestClinet(t *testing.T) {
	client := new(Clinet)
	macBookMachine := new(MacBook)

	client.InsetSquareUsbInComputer(macBookMachine)

	windowsMachine := new(Windows)
	windowsApdpater := &WindowPortAdpater{Windows: windowsMachine}
	client.InsetSquareUsbInComputer(windowsApdpater)
}
