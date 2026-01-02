package usb

/*
#cgo linux LDFLAGS: -L/usr/lib -lusb-1.0 -g
#cgo linux CFLAGS: -I/usr/include -I. -g -O0 -Wall
#cgo linux pkg-config: libusb-1.0

#cgo darwin LDFLAGS: -L/opt/homebrew/lib -lusb-1.0
#cgo darwin CFLAGS: -I/opt/homebrew/include -I/opt/homebrew/include/libusb-1.0  -I.

#include <libusb.h>
#include <devices.h>
#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type USBService struct {
	initialized bool
}

func NewInitializedUSBService() (*USBService, error) {
	i := C.libusb_init(nil)
	if i != 0 {
		return nil, errors.New("failed to initialize libusb")
	}

	return &USBService{
		initialized: true,
	}, nil
}

func (u *USBService) Initialized() bool {
	return u.initialized
}

func (u *USBService) ListDevices() ([]*USBDevice, error) {
	if !u.Initialized() {
		return nil, errors.New("libusb not initialized")
	}

	usbDevices := make([]*USBDevice, 0)

	var devices **C.struct_libusb_device

	// defer this to free the device list
	defer C.libusb_free_device_list(devices, C.int(1))

	// get the device list
	deviceCount := C.libusb_get_device_list(nil, &devices)

	for _, device := range unsafe.Slice(devices, deviceCount) {
		var handle *C.struct_libusb_device_handle
		ret := C.libusb_open(device, &handle)
		if ret != 0 {
			return nil, fmt.Errorf("failed to open device: %s", C.GoString(C.libusb_strerror(ret)))
		}

		var descriptor C.struct_libusb_device_descriptor
		ret = C.libusb_get_device_descriptor(device, &descriptor)
		if ret != 0 {
			return nil, fmt.Errorf("failed to get device descriptor: %s", C.GoString(C.libusb_strerror(ret)))
		}

		sn, err := u.stringValue(handle, descriptor.iSerialNumber)
		if err != nil {
			return nil, err
		}

		man, err := u.stringValue(handle, descriptor.iManufacturer)
		if err != nil {
			return nil, err
		}

		// var endpoint C.struct_libusb_endpoint_descriptor
		// var comp *C.struct_libusb_ss_endpoint_companion_descriptor
		// ret = C.libusb_get_ss_endpoint_companion_descriptor(nil, &endpoint, &comp)
		// if ret != 0 {
		// 	return nil, fmt.Errorf("failed to get endpoint descriptor: %s", C.GoString(C.libusb_strerror(ret)))
		// }

		ud := &USBDevice{
			busNumber:    uint8(C.libusb_get_bus_number(device)),
			address:      uint8(C.libusb_get_device_address(device)),
			speed:        int(C.libusb_get_device_speed(device)),
			port:         uint8(C.libusb_get_port_number(device)),
			device:       device,
			serialNumber: sn,
			manufacturer: man,
			// endpoint:     fmt.Sprintf("%x", uint8(endpoint.bEndpointAddress)),
		}

		usbDevices = append(usbDevices, ud)
	}

	return usbDevices, nil
}

func (u *USBService) stringValue(handle *C.struct_libusb_device_handle, idx C.uint8_t) (string, error) {
	data := (*C.uchar)(C.malloc(256))

	defer C.free(unsafe.Pointer(data))

	dataSize := C.libusb_get_string_descriptor_ascii(handle, idx, data, C.int(256))
	if dataSize < 0 {
		if dataSize == C.LIBUSB_ERROR_INVALID_PARAM {
			return "Not Provided", nil
		}
		return "", fmt.Errorf("failed to get string value: %s for index: %v", C.GoString(C.libusb_strerror(dataSize)), idx)
	}

	return string(C.GoBytes(unsafe.Pointer(data), dataSize)), nil
}
