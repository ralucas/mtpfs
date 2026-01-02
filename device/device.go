package device

import "github.com/ralucas/mtpfs/pkg/usb"

func List() ([]*usb.USBDevice, error) {
	u, err := usb.NewInitializedUSBService()
	if err != nil {
		return nil, err
	}

	return u.ListDevices()
}
