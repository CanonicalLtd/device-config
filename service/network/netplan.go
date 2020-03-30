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
	"bufio"
	"fmt"
	"github.com/CanonicalLtd/device-config/service/dbus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const netplanFilePath = "/etc/netplan/99-device-config.yaml"

// NetplanYAML defines the structure of the netplan YAML file
type NetplanYAML struct {
	Network Network `yaml:"network"`
}

// Network defines then network subsection of the netplan YAML
type Network struct {
	Version   int                 `yaml:"version"`
	Renderer  string              `yaml:"renderer,omitempty"`
	Ethernets map[string]Ethernet `yaml:"ethernets,omitempty"`
}

// Ethernet defines a single interface
type Ethernet struct {
	Use         bool                `yaml:"-"`
	Name        string              `yaml:"-"`
	DHCP4       string              `yaml:"dhcp4,omitempty"`
	Addresses   []string            `yaml:"addresses,omitempty"`
	NameServers map[string][]string `yaml:"nameservers,omitempty"`
	Gateway4    string              `yaml:"gateway4,omitempty"`
}

// NetplanService is the interface for the netplan service
type NetplanService interface {
	Apply() error
	Current() *NetplanYAML
	Store(ethernet Ethernet) error
}

// Netplan implements actions for managing netplan
type Netplan struct {
	deviceNetplan *NetplanYAML
	dBus          dbus.Service
}

// NewNetplan creates a netplan object from a config file
func NewNetplan(dBus dbus.Service) *Netplan {
	fmt.Println("Using netplan for network configuration")
	deviceNetplan := &NetplanYAML{Network: Network{Version: 2, Renderer: "networkd"}}

	data, err := readNetplanFile()
	if err != nil {
		// Cannot find the file, so set up an empty structure and store it
		srv := &Netplan{defaultNetplan(), dBus}
		_ = serializeNetplan(srv.deviceNetplan)
		return srv
	}

	if err := yaml.Unmarshal(data, deviceNetplan); err != nil {
		log.Printf("Error parsing the netplan file: %v", err)
		return &Netplan{defaultNetplan(), dBus}
	}
	return &Netplan{deviceNetplan, dBus}
}

func defaultNetplan() *NetplanYAML {
	log.Println("Using the default netplan configuration")
	return &NetplanYAML{Network: Network{Version: 2, Renderer: "networkd"}}
}

// Current returns the current netplan settings
func (np *Netplan) Current() *NetplanYAML {
	return np.deviceNetplan
}

// Apply applies the netplan configuration using dbus
func (np *Netplan) Apply() error {
	return np.dBus.NetplanApply()
}

// Store stores the updated network settings
func (np *Netplan) Store(ethernet Ethernet) error {
	// Validate the network addresses (for manual config only)
	if ethernet.DHCP4 != "true" {
		addresses := []string{}
		for _, a := range ethernet.Addresses {
			ip, mask, err := validateAddress(a)
			if err != nil {
				return err
			}
			addr := fmt.Sprintf("%s/%d", ip, mask)

			addresses = append(addresses, addr)
		}
		ethernet.Addresses = addresses
	}

	// Initialize the list of interfaces
	if np.deviceNetplan.Network.Ethernets == nil {
		np.deviceNetplan.Network.Ethernets = map[string]Ethernet{}
	}

	if !ethernet.Use {
		// Remove the configuration if it is not to be used
		delete(np.deviceNetplan.Network.Ethernets, ethernet.Name)
	} else {
		// Update the configuration for the interface
		np.deviceNetplan.Network.Ethernets[ethernet.Name] = ethernet
	}

	return serializeNetplan(np.deviceNetplan)
}

// readNetplanFile reads the current netplan file
var readNetplanFile = func() ([]byte, error) {
	return ioutil.ReadFile(netplanFilePath)
}

var writeNetplan = func(data []byte) error {
	// Write the YAML to the config file
	f, err := os.Create(netplanFilePath)
	if err != nil {
		return err
	}
	w := bufio.NewWriter(f)
	if _, err := w.Write(data); err != nil {
		return err
	}
	_ = w.Flush()

	return nil
}
