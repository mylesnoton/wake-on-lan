package main

import (
	"errors"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

// New wake on lan request
func New(macAddr string, bCast string) {
	macAddrArray, err := cleanAndConvertMacAddr(macAddr)
	handleErr(err)

	magicPacket, err := buildMagicPacket(macAddrArray)
	handleErr(err)

	send(bCast, magicPacket)
}

func cleanAndConvertMacAddr(macAddr string) ([]string, error) {
	matched, err := regexp.MatchString(`([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`, macAddr)
	handleErr(err)

	if matched != true {
		return nil, errors.New("MAC Address is not valid")
	}

	macAddr = strings.ToLower(macAddr)
	macAddr = strings.ReplaceAll(macAddr, ":", "-")

	macAddrSplit := strings.Split(macAddr, "-")

	return macAddrSplit, nil
}

func send(bCast string, magicPacket []byte) {
	bCastAddr, err := net.ResolveUDPAddr("udp4", bCast+":9")
	handleErr(err)

	conn, err := net.DialUDP("udp4", nil, bCastAddr)
	handleErr(err)

	defer conn.Close()

	_, err = conn.Write(magicPacket)
	handleErr(err)

	color.Green.Println("Broadcast " + strconv.Itoa(len(magicPacket)) + " bytes")
}

func handleErr(err error) {
	if err != nil {
		color.Red.Println(err.Error())
		return
	}
}
