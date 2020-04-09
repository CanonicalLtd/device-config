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
	"github.com/godbus/dbus"
	"testing"
	"time"
)

// mockBusObject mocks the dbus bus object interface
type mockBusObject struct {
}

// Call mocks a dbus method call
func (m *mockBusObject) Call(method string, flags dbus.Flags, args ...interface{}) *dbus.Call {
	c := dbus.Call{Err: nil}

	switch method {
	case "org.freedesktop.NetworkManager.GetAllDevices":
		c.Body = []interface{}{[]string{"/org/freedesktop/NetworkManager/Device/1"}}
	case "org.freedesktop.NetworkManager.Settings.Connection.GetSettings":
		c.Body = []interface{}{map[string]map[string]dbus.Variant{
			"ipv4":                     {"method": dbus.MakeVariant("manual")},
			"802-11-wireless":          {"ssid": dbus.MakeVariant("AirportWifi")},
			"802-11-wireless-security": {"psk": dbus.MakeVariant("TrustMe!")},
		}}
	}
	return &c
}

// Go mocks a dbus asynchronous call
func (m *mockBusObject) Go(method string, flags dbus.Flags, ch chan *dbus.Call, args ...interface{}) *dbus.Call {
	return nil
}

// GetProperty mocks dbus property retrieval
func (m *mockBusObject) GetProperty(p string) (dbus.Variant, error) {
	switch p {
	case "org.freedesktop.timedate1.Timezone":
		return dbus.MakeVariant("Europe/London"), nil
	case "org.freedesktop.timedate1.NTP":
		return dbus.MakeVariant(false), nil
	case "org.freedesktop.timedate1.TimeUSec":
		return dbus.MakeVariant(uint64(time.Now().Unix() * 1e6)), nil
	case "org.freedesktop.NetworkManager.Device.State":
		return dbus.MakeVariant(uint32(100)), nil
	case "org.freedesktop.NetworkManager.Device.Ip4Config":
		return dbus.MakeVariant("/org/freedesktop/NetworkManager/Ip4Config/1"), nil
	case "org.freedesktop.NetworkManager.IP4Config.Nameservers":
		return dbus.MakeVariant([]uint32{16885952}), nil
	case "org.freedesktop.NetworkManager.IP4Config.AddressData":
		return dbus.MakeVariant([]map[string]dbus.Variant{{"address": dbus.MakeVariant("192.168.1.1"), "prefix": dbus.MakeVariant(uint32(24))}}), nil
	case "org.freedesktop.NetworkManager.Device.Dhcp4Config":
		return dbus.MakeVariant("/org/freedesktop/NetworkManager/Dhcp4Config/1"), nil
	case "org.freedesktop.NetworkManager.Device.DeviceType":
		return dbus.MakeVariant(uint32(2)), nil
	}
	return dbus.MakeVariant(""), nil
}

// Destination mocks a dbus destination call
func (m *mockBusObject) Destination() string {
	return ""
}

// Path mocks a dbus path call
func (m *mockBusObject) Path() dbus.ObjectPath {
	return ""
}

func TestDBus_NetplanApply(t *testing.T) {
	busObject = getMockBusObject
	db := &DBus{systemBus: nil}

	if err := db.NetplanApply(); err != nil {
		t.Errorf("NetplanApply() error = %v", err)
	}
}
