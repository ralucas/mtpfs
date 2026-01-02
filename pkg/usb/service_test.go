package usb_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ralucas/mtpfs/pkg/usb"
)

func TestListDevices(t *testing.T) {
	u, err := usb.NewInitializedUSBService()
	require.NoError(t, err)

	devices, err := u.ListDevices()

	for _, d := range devices {
		fmt.Printf("%+v\n", d)
	}

	assert.NoError(t, err)
	assert.True(t, len(devices) > 0)
}
