package main

import (
	"encoding/hex"
	"errors"
)

func buildMagicPacket(macAddr []string) ([]byte, error) {
	if len(macAddr) != 6 {
		return nil, errors.New("MAC Address is not valid")
	}

	magicPacket := make([]byte, 102)

	for i := 0; i < 6; i++ {
		magicPacket[i] = 255
	}

	position := 6

	for i := 0; i < 16; i++ {
		for j := 0; j < 6; j++ {
			macBytes, err := hex.DecodeString(macAddr[j])
			handleErr(err)

			magicPacket[position] = byte(macBytes[0])
			position++
		}
	}

	return magicPacket, nil
}
