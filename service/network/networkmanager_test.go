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
	"github.com/CanonicalLtd/device-config/config"
	"github.com/CanonicalLtd/device-config/service/dbus"
	"reflect"
	"testing"
)

func mockInterfacesValid() ([]HardwareInterface, error) {
	return []HardwareInterface{
		{Name: "eth0", MACAddress: "eth0-mac-address"},
	}, nil
}

func mockInterfacesError() ([]HardwareInterface, error) {
	return nil, fmt.Errorf("MOCK interfaces error")
}

func TestNetManager_Current(t *testing.T) {
	settings := config.DefaultArgs()
	settings.UseNetworkManager = true
	defaultNetplan := &NetplanYAML{
		Network: Network{
			Version: 2, Renderer: "NetworkManager",
			Ethernets: map[string]Ethernet{
				"eth0": {Name: "eth0", NameServers: map[string][]string{"addresses": {"192.168.1.1", "8.8.8.8"}}, Addresses: []string{"192.168.1.100/24"}, Use: true},
			},
		},
	}
	defaultNetplanDefault := &NetplanYAML{Network: Network{Version: 2, Renderer: "NetworkManager", Ethernets: map[string]Ethernet{}}}

	tests := []struct {
		name       string
		interfaces func() ([]HardwareInterface, error)
		want       *NetplanYAML
	}{
		{"valid-default", mockInterfacesValid, defaultNetplan},
		{"invalid-interfaces", mockInterfacesError, defaultNetplanDefault},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interfaces = tt.interfaces
			np := Factory(settings, &dbus.MockDbus{})
			got := np.Current()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Current() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNetManager_Store(t *testing.T) {
	settings := config.DefaultArgs()
	settings.UseNetworkManager = true
	validDHCP := Ethernet{DHCP4: "true"}
	validManual := Ethernet{DHCP4: "", Addresses: []string{"192.168.1.100"}, NameServers: map[string][]string{"addresses": {"192.168.1.1"}}, Gateway4: "192.168.1.1"}

	tests := []struct {
		name     string
		ethernet Ethernet
		wantErr  bool
	}{
		{"valid-dhcp", validDHCP, false},
		{"valid-manual", validManual, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interfaces = mockInterfacesValid
			np := Factory(settings, &dbus.MockDbus{})
			if err := np.Store(tt.ethernet); (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNetManager_Apply(t *testing.T) {
	settings := config.DefaultArgs()
	settings.UseNetworkManager = true
	np := Factory(settings, &dbus.MockDbus{})
	if err := np.Apply(); err != nil {
		t.Errorf("Apply() error = %v, wantErr %v", err, nil)
	}
}
