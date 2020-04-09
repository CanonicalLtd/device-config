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
	"testing"
)

func TestDBus_NMIsRunning(t *testing.T) {
	busObject = getMockBusObject
	db := &DBus{systemBus: nil}

	if err := db.NMIsRunning(); err != nil {
		t.Errorf("NMIsRunning() error = %v", err)
	}
}

func TestDBus_NMDevices(t *testing.T) {
	busObject = getMockBusObject
	db := &DBus{systemBus: nil}

	_, err := db.NMDevices()
	if err != nil {
		t.Errorf("NMDevices() error = %v", err)
	}
}

func TestDBus_NMInterfaceConfig(t *testing.T) {
	busObject = getMockBusObject
	db := &DBus{systemBus: nil}

	got := db.NMInterfaceConfig("/org/freedesktop/NetworkManager/Device/1")
	if len(got.NameServers) != 1 && got.NameServers[0] != "192.168.1.1" {
		t.Errorf("NMInterfaceConfig() nameservers = %v, want %v", got.NameServers, "[192.168.1.1]")
	}
}

func TestDBus_NMInterfaceConfigUpdate(t *testing.T) {
	dhcpManual := NMDeviceSettings{
		DHCP4:       false,
		AddressData: []NMDeviceAddress{{Address: "192.168.2.100", Prefix: 24}},
		NameServers: []string{"192.168.2.1", "8.8.8.8"},
		Gateway:     "192.168.2.1",
	}
	dhcpAuto := NMDeviceSettings{DHCP4: true, IsWifi: true, SSID: "AirportWifi", Password: "TrustMe!"}

	tests := []struct {
		name    string
		eth     NMDeviceSettings
		wantErr bool
	}{
		{"valid-manual", dhcpManual, false},
		{"valid-dhcp", dhcpAuto, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			busObject = getMockBusObject
			db := &DBus{systemBus: nil}

			if err := db.NMInterfaceConfigUpdate("/org/freedesktop/NetworkManager/Device/1", tt.eth); (err != nil) != tt.wantErr {
				t.Errorf("NMInterfaceConfigUpdate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
