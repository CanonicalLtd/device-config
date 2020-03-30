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
	"fmt"
	"github.com/CanonicalLtd/device-config/config"
	"github.com/CanonicalLtd/device-config/service/dbus"
	"net"
	"strings"
)

// Service is the interface for the netplan/network-manager services
type Service interface {
	Apply() error
	Current() *NetplanYAML
	Store(ethernet Ethernet) error
}

// HardwareInterface represents a hardware network interface
type HardwareInterface struct {
	Name       string
	MACAddress string
}

// Factory creates the relevant netplan or network manager service
func Factory(settings *config.Settings, dBus dbus.Service) Service {
	// Take over netplan config for this device
	_ = TakeOver()

	if settings.UseNetworkManager {
		return NewNetworkManager(dBus)
	}
	return NewNetplan(dBus)
}

// Interfaces fetches the list of network interfaces
var Interfaces = func() ([]HardwareInterface, error) {
	ifaces := []HardwareInterface{}
	interfaces, err := net.Interfaces()
	if err != nil {
		return ifaces, err
	}

	for _, i := range interfaces {
		// Select the real network interfaces only
		if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
			ifaces = append(ifaces, HardwareInterface{i.Name, i.HardwareAddr.String()})
		}
	}
	return ifaces, nil
}

// validateAddress check we have a valid network address e.g. 192.168.1.111/24
func validateAddress(addr string) (string, uint32, error) {
	// Check if we have a full network address
	ip, ipNet, err := net.ParseCIDR(addr)
	if err == nil {
		ones, _ := ipNet.Mask.Size()
		return ip.String(), uint32(ones), nil
	}

	// Check if we have a slash e.g. 192.168.1.111/255.255.255.0
	parts := strings.Split(addr, "/")

	// Check if we just use the IP
	ip = net.ParseIP(parts[0])
	if ip == nil {
		return "", 0, fmt.Errorf("the network address and mask are invalid")
	}

	// Calculate the default mask
	ones, _ := ip.DefaultMask().Size()
	return ip.String(), uint32(ones), nil
}
