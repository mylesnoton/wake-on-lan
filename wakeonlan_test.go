package main

import (
	"testing"
)

func TestCleanAndConvertMacAddr(t *testing.T) {
	macAddr, err := cleanAndConvertMacAddr("00-80-C8-E3-4C-BD")

	if err != nil {
		t.Fail()
	}

	if len(macAddr) != 6 {
		t.Fail()
	}
}
func TestCleanAndConvertMacAddrInvalidShort(t *testing.T) {
	_, err := cleanAndConvertMacAddr("00-80-C8-E3-4C")

	if err == nil {
		t.Fail()
	}
}
