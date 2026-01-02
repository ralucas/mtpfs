package usb

type USBDevice struct {
	busNumber    uint8
	address      uint8
	speed        int
	port         uint8
	device       *_Ctype_struct_libusb_device
	serialNumber string
	manufacturer string
	endpoint     string
}

func (d *USBDevice) BusNumber() uint8 {
	return d.busNumber
}

func (d *USBDevice) Address() uint8 {
	return d.address
}

func (d *USBDevice) Speed() int {
	return d.speed
}

func (d *USBDevice) Port() uint8 {
	return d.port
}

func (d *USBDevice) Device() *_Ctype_struct_libusb_device {
	return d.device
}
