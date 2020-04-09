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
)

const (
	wifiType = uint32(2)
)

// NMDeviceAddress holds the details of a network address
type NMDeviceAddress struct {
	Address string
	Prefix  uint32
}

// NMDeviceSettings holds the configuration for a NM device
type NMDeviceSettings struct {
	DHCP4       bool
	State       uint32
	AddressData []NMDeviceAddress
	NameServers []string
	Gateway     string
	IsWifi      bool
	SSID        string
	Password    string
}

// NMIsRunning checks if the network manager service is running
func (db *DBus) NMIsRunning() error {
	netman := db.getBusObject("org.freedesktop.DBus", "/org/freedesktop/DBus")
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
		netman1 := db.getBusObject("org.freedesktop.NetworkManager", dev)

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
func (db *DBus) NMInterfaceConfig(p string) *NMDeviceSettings {
	eth := NMDeviceSettings{}
	netman := db.getBusObject("org.freedesktop.NetworkManager", p)

	// Check if the interface is used
	state, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.State")
	if err == nil {
		eth.State = state.Value().(uint32)
	}

	// Ipv4 config
	ip4, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.Ip4Config")
	if err == nil {
		p := fmt.Sprintf("%v", ip4.Value())
		eth = db.ip4Config(p, eth)
	}

	// DHCP config
	settings, err := db.connectionConfig(p)
	if err == nil {
		if settings["ipv4"]["method"].Value().(string) == "auto" {
			eth.DHCP4 = true
		}
	}

	// Check for wifi device
	deviceType, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.DeviceType")
	if err != nil {
		return &eth
	}
	if deviceType.Value() != wifiType {
		return &eth
	}

	// This is a wifi type
	eth.IsWifi = true
	ssid, ok := settings["802-11-wireless"]["ssid"]
	if ok {
		eth.SSID = ssid.String()
	}
	pwd, ok := settings["802-11-wireless-security"]["psk"]
	if ok {
		eth.Password = pwd.String()
	}

	return &eth
}

// NMInterfaceConfigUpdate stores the updated configuration for a hardware interface
func (db *DBus) NMInterfaceConfigUpdate(p string, eth NMDeviceSettings) error {
	// Generate the dbus settings for device
	settings := db.createSettings(eth)

	// Get the device object
	netman := db.getBusObject("org.freedesktop.NetworkManager", p)

	// Get the connection for the device
	activeConn, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.ActiveConnection")
	if err != nil {
		// Connection does not exist: add a new activeConn for the interface
		return db.addAndActivateConnection(p, settings)
	}

	// Get the connection path from the active connection
	activePath := fmt.Sprintf("%v", activeConn.Value())
	connection := db.getBusObject("org.freedesktop.NetworkManager", activePath)
	conn, err := connection.GetProperty("org.freedesktop.NetworkManager.Connection.Active.Connection")
	if err != nil {
		return err
	}

	// Update the connection
	connPath := fmt.Sprintf("%v", conn.Value())
	return db.updateConnection(connPath, settings)
}

func (db *DBus) connectionConfig(p string) (map[string]map[string]dbus.Variant, error) {
	netman := db.getBusObject("org.freedesktop.NetworkManager", p)

	// Get the active connection details
	activeConn, err := netman.GetProperty("org.freedesktop.NetworkManager.Device.ActiveConnection")
	if err != nil {
		return nil, err
	}
	activePath := fmt.Sprintf("%v", activeConn.Value())
	active := db.getBusObject("org.freedesktop.NetworkManager", activePath)

	// Get the connection settings
	conn, err := active.GetProperty("org.freedesktop.NetworkManager.Connection.Active.Connection")
	if err != nil {
		return nil, err
	}
	connPath := fmt.Sprintf("%v", conn.Value())

	// Get the current settings for the connection
	netmanConn := db.getBusObject("org.freedesktop.NetworkManager", connPath)

	var s map[string]map[string]dbus.Variant
	err = netmanConn.Call("org.freedesktop.NetworkManager.Settings.Connection.GetSettings", 0).Store(&s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (db *DBus) createSettings(eth NMDeviceSettings) map[string]map[string]dbus.Variant {
	settings := map[string]map[string]dbus.Variant{
		"ipv4": {},
	}

	// Wifi connection details
	if eth.IsWifi {
		settings["802-11-wireless"] = map[string]dbus.Variant{}
		settings["802-11-wireless-security"] = map[string]dbus.Variant{}
		settings["802-11-wireless"]["ssid"] = dbus.MakeVariant(eth.SSID)
		settings["802-11-wireless-security"]["key-mgmt"] = dbus.MakeVariant("wpa-psk")
		settings["802-11-wireless-security"]["psk"] = dbus.MakeVariant(eth.Password)
	}

	if eth.DHCP4 {
		// DHCP network config
		settings["ipv4"]["method"] = dbus.MakeVariant("auto")
		return settings
	}

	// Manual network config
	settings["ipv4"]["method"] = dbus.MakeVariant("manual")
	settings["ipv4"]["gateway"] = dbus.MakeVariant(eth.Gateway)

	// Address
	addressData := []map[string]dbus.Variant{}
	for _, a := range eth.AddressData {
		address := map[string]dbus.Variant{
			"address": dbus.MakeVariant(a.Address),
			"prefix":  dbus.MakeVariant(a.Prefix),
		}
		addressData = append(addressData, address)
	}
	settings["ipv4"]["address-data"] = dbus.MakeVariant(addressData)

	// Name servers
	dns := []uint32{}
	for _, n := range eth.NameServers {
		dns = append(dns, ipToUint32(n))
	}
	settings["ipv4"]["dns"] = dbus.MakeVariant(dns)

	return settings
}

// addAndActivateConnection adds a new connection for an unconfigured device
func (db *DBus) addAndActivateConnection(devicePath string, settings map[string]map[string]dbus.Variant) error {
	netman := db.getBusObject("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")

	var pathConnection, pathActive string
	if err := netman.Call("org.freedesktop.NetworkManager.AddAndActivateConnection", 0, settings, dbus.ObjectPath(devicePath), dbus.ObjectPath("/")).Store(&pathConnection, &pathActive); err != nil {
		return err
	}
	return nil
}

// updateConnection updates the configuration for a device
func (db *DBus) updateConnection(connPath string, settings map[string]map[string]dbus.Variant) error {
	netman := db.getBusObject("org.freedesktop.NetworkManager", connPath)

	// Get the current settings for the connection
	var s map[string]map[string]dbus.Variant
	err := netman.Call("org.freedesktop.NetworkManager.Settings.Connection.GetSettings", 0).Store(&s)
	if err != nil {
		return err
	}

	// Setting the connection ID details and update the connection
	settings["connection"] = s["connection"]
	call := netman.Call("org.freedesktop.NetworkManager.Settings.Connection.Update", 0, settings)
	return call.Err
}

// ip4Config decodes the IPv4 config object
func (db *DBus) ip4Config(p string, eth NMDeviceSettings) NMDeviceSettings {
	netman := db.getBusObject("org.freedesktop.NetworkManager", p)

	gateway, err := netman.GetProperty("org.freedesktop.NetworkManager.IP4Config.Gateway")
	if err == nil {
		eth.Gateway = gateway.Value().(string)
	}

	nameServers, err := netman.GetProperty("org.freedesktop.NetworkManager.IP4Config.Nameservers")
	if err == nil {
		ns := []string{}
		for _, n := range nameServers.Value().([]uint32) {
			ns = append(ns, uint32ToIP(n))
		}
		eth.NameServers = ns
	}

	// Decode the addresses property
	address, err := netman.GetProperty("org.freedesktop.NetworkManager.IP4Config.AddressData")
	if err == nil {
		addresses := []NMDeviceAddress{}
		aa := address.Value().([]map[string]dbus.Variant)
		for _, a := range aa {
			addr := NMDeviceAddress{Address: a["address"].Value().(string), Prefix: a["prefix"].Value().(uint32)}
			addresses = append(addresses, addr)
		}
		eth.AddressData = addresses
	}
	return eth
}

// getDevices gets the paths of the network interfaces
func (db *DBus) getDevices() ([]string, error) {
	netman := db.getBusObject("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
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

func ipToUint32(ip string) uint32 {
	netIP := net.ParseIP(ip)
	return binary.LittleEndian.Uint32(netIP.To4())
}
