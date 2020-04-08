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

package web

import (
	"fmt"
	"github.com/CanonicalLtd/device-config/datastore"
	"github.com/CanonicalLtd/device-config/service"
	"github.com/CanonicalLtd/device-config/service/network"
	"github.com/CanonicalLtd/device-config/service/snapd"
	"github.com/snapcore/snapd/client"
	"time"
)

type mockAuth struct{}

func (auth *mockAuth) ValidateSession(username, sessionID string) (*datastore.Session, error) {
	if username == "invalid" {
		return nil, fmt.Errorf("MOCK error validating session")
	}
	return &datastore.Session{
		ID:        1,
		Username:  "generated-user",
		SessionID: "generated-session-id",
		Expires:   time.Now().Add(24 * time.Hour),
	}, nil
}
func (auth *mockAuth) CreateSession(token string, expires time.Time) (*datastore.Session, error) {
	if token == "invalid" {
		return nil, fmt.Errorf("MOCK error creating session")
	}
	return &datastore.Session{
		ID:        1,
		Username:  "generated-user",
		SessionID: "generated-session-id",
		Expires:   expires,
	}, nil
}

type mockNetplan struct {
	applyError bool
}

func (np *mockNetplan) Apply() error {
	if np.applyError {
		return fmt.Errorf("MOCK apply error")
	}
	return nil
}

func (np *mockNetplan) Current() *network.NetplanYAML {
	return &network.NetplanYAML{
		Network: network.Network{
			Version:  2,
			Renderer: "networkd",
			Ethernets: map[string]network.Ethernet{
				"enp3s0": {DHCP4: "true"},
				"eth1":   {Addresses: []string{"192.168.2.200/192.168.2.255"}, Gateway4: "192.168.2.1", NameServers: map[string][]string{"addresses": {"8.8.8.8"}}},
			},
		},
	}
}

func (np *mockNetplan) Store(ethernet network.Ethernet) error {
	if ethernet.Name == "invalid" {
		return fmt.Errorf("MOCK store error")
	}
	return nil
}

type mockSnapd struct {
	confError     bool
	setConfError  bool
	servicesError bool
	snapsErr      bool
}

func (snap *mockSnapd) Conf(name string) (map[string]interface{}, error) {
	if snap.confError {
		return nil, fmt.Errorf("MOCK snapd conf error")
	}
	httpProxy := map[string]string{"http": "192.168.1.100:4000"}
	return map[string]interface{}{"proxy": httpProxy}, nil
}

func (snap *mockSnapd) SetConf(name string, patch map[string]interface{}) (string, error) {
	if snap.setConfError {
		return "", fmt.Errorf("MOCK snapd setconf error")
	}
	return "100", nil
}

func (snap *mockSnapd) AppServices(names []string) ([]*client.AppInfo, error) {
	if snap.servicesError {
		return nil, fmt.Errorf("MOCK snap services error")
	}
	return []*client.AppInfo{
		{Snap: "chuck-norris-webserver", Name: "daemon", Enabled: true, Active: true},
		{Snap: "super-agent", Name: "service", Enabled: true, Active: false},
	}, nil
}

func (snap *mockSnapd) List(names []string, opts *client.ListOptions) ([]snapd.Snap, error) {
	if snap.snapsErr {
		return nil, fmt.Errorf("MOCK snap list error")
	}
	return []snapd.Snap{}, nil
}

type mockTime struct{}

func (t *mockTime) Current() *service.Time {
	return &service.Time{
		Timezones: []string{"America/Barbados", "America/Hawaii"},
		Timezone:  "America/Barbados",
		NTP:       true,
		Time:      time.Now(),
	}
}

func (t *mockTime) Apply(ntp bool, timezone, setTime string) error {
	if timezone == "invalid" {
		return fmt.Errorf("MOCK time apply error")
	}
	return nil
}

func mockInterfacesValid() ([]network.HardwareInterface, error) {
	return []network.HardwareInterface{
		{Name: "enp3s0", MACAddress: "enp3s0-mac-address"},
		{Name: "eth0", MACAddress: "eth0-mac-address"},
		{Name: "eth1", MACAddress: "eth1-mac-address"},
	}, nil
}

func mockInterfacesNone() ([]network.HardwareInterface, error) {
	return []network.HardwareInterface{}, fmt.Errorf("MOCK error")
}
