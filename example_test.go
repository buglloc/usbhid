package usbhid_test

import (
	"fmt"
	"log"

	"github.com/buglloc/usbhid"
)

func ExampleEnumerate() {
	devices, err := usbhid.Enumerate() // no filtering
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		fmt.Printf("Device: 0x%04x:0x%04x\n", device.VendorId(), device.ProductId())
		fmt.Printf("\tManufacturer:  %s\n", device.Manufacturer())
		fmt.Printf("\tProduct:       %s\n", device.Product())
		fmt.Printf("\tSerial Number: %s\n", device.SerialNumber())
		fmt.Printf("\tUsage:         0x%04x/0x%04x\n", device.UsagePage(), device.Usage())
	}
}

func ExampleWithVidFilter() {
	devices, err := usbhid.Enumerate(usbhid.WithVidFilter(0x1050))
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices {
		fmt.Printf("Device: 0x%04x:0x%04x\n", device.VendorId(), device.ProductId())
		fmt.Printf("\tManufacturer:  %s\n", device.Manufacturer())
		fmt.Printf("\tProduct:       %s\n", device.Product())
		fmt.Printf("\tSerial Number: %s\n", device.SerialNumber())
		fmt.Printf("\tUsage:         0x%04x/0x%04x\n", device.UsagePage(), device.Usage())
	}
}

func ExampleWithDeviceFilterFunc() {
	device, err := usbhid.Get(
		usbhid.WithDeviceFilterFunc(func(d *usbhid.Device) bool {
			// filtering by free HID VID/PID from v-usb
			if d.VendorId() != 0x16c0 {
				return false
			}

			if d.ProductId() != 0x05df {
				return false
			}

			return true
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Device: 0x%04x:0x%04x\n", device.VendorId(), device.ProductId())
	fmt.Printf("\tManufacturer:  %s\n", device.Manufacturer())
	fmt.Printf("\tProduct:       %s\n", device.Product())
	fmt.Printf("\tSerial Number: %s\n", device.SerialNumber())
}
