package main

import (
	"fmt"
	"log"

	"github.com/buglloc/usbhid"
)

func main() {
	devices, err := usbhid.Enumerate(
		usbhid.WithVidFilter(0x1050),
	)
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
