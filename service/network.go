// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package service

import (
	"bytes"
	"net"
)

// NetworkInterface represents a hardware network interface
type NetworkInterface struct {
	Name       string
	MACAddress string
}

// Interfaces fetches the list of network interfaces
func Interfaces() ([]NetworkInterface, error) {
	ifaces := []NetworkInterface{}
	interfaces, err := net.Interfaces()
	if err != nil {
		return ifaces, err
	}

	for _, i := range interfaces {
		// Select the real network interfaces only
		if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
			ifaces = append(ifaces, NetworkInterface{i.Name, i.HardwareAddr.String()})
		}
	}
	return ifaces, nil
}

// ValidateIP4 checks that we have a valid IPv4 address
func ValidateIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
