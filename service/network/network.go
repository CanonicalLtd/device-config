/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package network

import (
	"bytes"
	"net"
)

// NetworkService is the interface for the netplan/network-manager services
type NetworkService interface {
	Apply() error
	Current() *NetplanYAML
	Store(ethernet Ethernet) error
}

// NetworkInterface represents a hardware network interface
type NetworkInterface struct {
	Name       string
	MACAddress string
}

// Interfaces fetches the list of network interfaces
var Interfaces = func() ([]NetworkInterface, error) {
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

// ValidateIP checks that we have a valid IPv4 address
func ValidateIP(ip string) bool {
	return net.ParseIP(ip) != nil
}
