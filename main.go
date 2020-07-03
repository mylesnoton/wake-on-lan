package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/gookit/color"
)

func main() {
	var macAddr string
	var bCast string

	flag.StringVar(&macAddr, "mac", "", "The MAC address of the machine you want to wake up")
	flag.StringVar(&macAddr, "bCast", "192.168.1.255", "The broadcast address on your network")
	flag.Parse()

	if macAddr == "" {
		fmt.Println("Enter MAC Address...")
		_, err := fmt.Scan(&macAddr)
		handleInputErr(err)
	}

	if len(macAddr) < 16 {
		handleInputErr(errors.New("Please enter a valid MAC Address, eg: XX:XX:XX:XX:XX:XX"))
	}

	New(macAddr, bCast)
}

func handleInputErr(err error) {
	color.Red.Println(err.Error)
	os.Exit(1)
}
