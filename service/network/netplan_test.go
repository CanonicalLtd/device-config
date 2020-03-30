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
	"fmt"
	"github.com/CanonicalLtd/device-config/service/dbus"
	"reflect"
	"testing"
)

var readNetplanFileError = func() ([]byte, error) {
	return nil, fmt.Errorf("MOCK error reading netplan file")
}
var readNetplanFileSuccess = func() ([]byte, error) {
	return []byte(`
network:
  version: 2
  renderer: networkd
  ethernets:
    enp3s0:
      dhcp4: true`), nil
}
var readNetplanFileUnused = func() ([]byte, error) {
	return []byte(`
network:
  version: 2
  renderer: networkd`), nil
}
var readNetplanFileBad = func() ([]byte, error) {
	return []byte("\u1000"), nil
}

func TestNetplan_Apply(t *testing.T) {
	// Mock the writing of the YAML file
	writeNetplan = func(data []byte) error {
		return nil
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"valid", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			np := NewNetplan(&dbus.MockDbus{})
			if err := np.Apply(); (err != nil) != tt.wantErr {
				t.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetplan_Current(t *testing.T) {
	// Mock the writing of the YAML file
	writeNetplan = func(data []byte) error {
		return nil
	}

	// Expected response from the mock
	var yaml = &NetplanYAML{Network: Network{
		Version:  2,
		Renderer: "networkd",
		Ethernets: map[string]Ethernet{
			"enp3s0": {DHCP4: "true"},
		},
	}}

	tests := []struct {
		name     string
		mockRead func() ([]byte, error)
		want     *NetplanYAML
	}{
		{"valid-default", readNetplanFileError, defaultNetplan()},
		{"valid-existing", readNetplanFileSuccess, yaml},
		{"invalid-existing", readNetplanFileBad, defaultNetplan()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readNetplanFile = tt.mockRead
			np := NewNetplan(&dbus.MockDbus{})
			if got := np.Current(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetplan_Store(t *testing.T) {
	// Mock the writing of the YAML file
	writeNetplan = func(data []byte) error {
		return nil
	}

	tests := []struct {
		name     string
		ethernet Ethernet
		mockRead func() ([]byte, error)
		wantErr  bool
	}{
		{"valid", Ethernet{Name: "eth0", DHCP4: "true", Use: true}, readNetplanFileError, false},
		{"valid-update", Ethernet{Name: "enp3s0", DHCP4: "true", Use: true}, readNetplanFileSuccess, false},
		{"valid-update-unused", Ethernet{Name: "enp3s0", DHCP4: "true", Use: false}, readNetplanFileSuccess, false},
		{"valid-manual", Ethernet{Name: "eth0", DHCP4: "", Addresses: []string{"192.168.1.100/24"}, Use: true}, readNetplanFileError, false},
		{"valid-manual-update", Ethernet{Name: "eth0", DHCP4: "", Addresses: []string{"192.168.1.100/24"}, Use: true}, readNetplanFileUnused, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the reading of the netplan file
			readNetplanFile = tt.mockRead

			np := NewNetplan(&dbus.MockDbus{})
			if err := np.Store(tt.ethernet); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
