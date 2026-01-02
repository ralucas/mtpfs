#ifndef DEVICES_H
#define DEVICES_H

#include <libusb.h>
#include <stdio.h>
#include <stdlib.h>
#include <stddef.h>

libusb_device** list_libusb_devices() {
	libusb_init(NULL);
	libusb_device **list;
	ssize_t cnt = libusb_get_device_list(NULL, &list);
	printf("Number of USB devices found: %ld\n", cnt);
	if (cnt < 0) {
		perror("no devices found");
	}
	return list;
}

#endif