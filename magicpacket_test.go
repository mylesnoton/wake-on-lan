package main

import (
	"strings"
	"testing"
)

func TestBuildMagicPacketLength(t *testing.T) {
	macAddrSplit := strings.Split("00-80-C8-E3-4C-BD", "-")
	theMagicPacket, _ := buildMagicPacket(macAddrSplit)

	if len(theMagicPacket) != 102 {
		t.Fail()
	}
}

func TestBuildMagicPacketStartingBytes(t *testing.T) {
	macAddrSplit := strings.Split("00-80-C8-E3-4C-BD", "-")
	theMagicPacket, _ := buildMagicPacket(macAddrSplit)

	if theMagicPacket[0] != 255 || theMagicPacket[5] != 255 {
		t.Fail()
	}
}

func TestBuildMagicPacketMacBytes(t *testing.T) {
	macAddrSplit := strings.Split("00-80-C8-E3-4C-BD", "-")
	theMagicPacket, _ := buildMagicPacket(macAddrSplit)

	if theMagicPacket[6] != 0 || theMagicPacket[11] != 189 {
		t.Fail()
	}

	if theMagicPacket[12] != 0 || theMagicPacket[17] != 189 {
		t.Fail()
	}

	if theMagicPacket[96] != 0 || theMagicPacket[101] != 189 {
		t.Fail()
	}
}

func TestBuildMagicPacketInvalidMacAddr(t *testing.T) {
	macAddrSplit := strings.Split("00-80-C8-E3-4C", "-")
	_, err := buildMagicPacket(macAddrSplit)

	if err == nil {
		t.Fail()
	}
}
