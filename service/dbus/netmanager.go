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

package dbus

import (
	"encoding/binary"
	"fmt"
	"github.com/godbus/dbus"
	"net"
	"strings"
)

// NMIsRunning checks if the network manager service is running
func (db *DBus) NMIsRunning() error {
	netman := db.systemBus.Object("org.freedesktop.DBus", "/org/freedesktop/DBus")
	call := netman.Call("org.freedesktop.DBus.GetNameOwner", 0, "org.freedesktop.NetworkManager", false)
	return call.Err
}

// NMDevices gets the paths of the network interfaces
func (db *DBus) NMDevices() (map[string]string, error) {
	// Get the object paths of the devices
	devices, err := db.getDevices()
	if err != nil {
		return nil, err
	}

	ifaces := map[string]string{}

	// Get the interface name for the device paths
	for _, dev := range devices {
		netman1 := db.systemBus.Object("org.freedesktop.NetworkManager", dbus.ObjectPath(dev))

		// Interface name
		iface, err := netman1.GetProperty("org.freedesktop.NetworkManager.Device.Interface")
		if err != nil {
			return nil, err
		}
		ifaces[iface.Value().(string)] = dev
	}

	return ifaces, nil
}

// NMInterfaceConfig gets the details of the active interfaces
func (db *DBus) NMInterfaceConfig(p string) map[string]string {
	eth := map[string]string{}
	netman := db.systemBus.Object("org.freedesktop.NetworkManager", dbus.ObjectPath(p))

	// Check if the interface is used
	state, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.State")
	if err != nil || state.String() < "100" {
		eth["use"] = "false"
		return eth
	}
	eth["use"] = "true"

	// Ipv4 config
	ip4, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.Ip4Config")
	if err == nil {
		p := fmt.Sprintf("%v", ip4.Value())
		eth = db.ip4Config(p, eth)
	}

	dhcp4, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.Dhcp4Config")
	if err == nil {
		p := fmt.Sprintf("%v", dhcp4.Value())
		if len(p) > 0 {
			eth["dhcp4"] = "true"
		}
	}

	return eth
}

// ip4Config decodes the IPv4 config object
func (db *DBus) ip4Config(p string, eth map[string]string) map[string]string {
	netman := db.systemBus.Object("org.freedesktop.NetworkManager", dbus.ObjectPath(p))

	gateway, err := netman.GetProperty("org.freedesktop.NetworkManager.IP4Config.Gateway")
	if err == nil {
		eth["gateway"] = gateway.Value().(string)
	}

	nameServers, err := netman.GetProperty("org.freedesktop.NetworkManager.IP4Config.Nameservers")
	if err == nil {
		ns := []string{}
		for _, n := range nameServers.Value().([]uint32) {
			ns = append(ns, uint32ToIP(n))
		}
		eth["nameservers"] = strings.Join(ns, ",")
	}

	// Decode the addresses property
	address, err := netman.GetProperty("org.freedesktop.NetworkManager.IP4Config.AddressData")
	if err == nil {
		addresses := []string{}
		aa := address.Value().([]map[string]dbus.Variant)
		for _, a := range aa {
			addr := fmt.Sprintf("%v/%v", a["address"].Value(), a["prefix"].Value())
			addresses = append(addresses, addr)
		}

		eth["addresses"] = strings.Join(addresses, ",")
	}
	return eth
}

// getDevices gets the paths of the network interfaces
func (db *DBus) getDevices() ([]string, error) {
	netman := db.systemBus.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	var devices []string
	err := netman.Call("org.freedesktop.NetworkManager.GetAllDevices", 0).Store(&devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func uint32ToIP(ipInt uint32) string {
	var ipByte [4]byte
	ip := ipByte[:]
	binary.LittleEndian.PutUint32(ip, ipInt)
	return net.IPv4(ip[0], ip[1], ip[2], ip[3]).String()
}
