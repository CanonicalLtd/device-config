// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package service

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

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
}

// NewNetplan creates a netplan object from a config file
func NewNetplan(path string) *Netplan {
	deviceNetplan := &NetplanYAML{Network: Network{Version: 2, Renderer: "networkd"}}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		// Cannot find the file, so set up an empty structure
		return &Netplan{defaultNetplan()}
	}

	if err := yaml.Unmarshal(data, deviceNetplan); err != nil {
		log.Printf("Error parsing the netplan file: %v", err)
		return &Netplan{defaultNetplan()}
	}
	return &Netplan{deviceNetplan}
}

func defaultNetplan() *NetplanYAML {
	return &NetplanYAML{Network: Network{Version: 2, Renderer: "networkd"}}
}

// Current returns the current netplan settings
func (np *Netplan) Current() *NetplanYAML {
	return np.deviceNetplan
}

// Apply applies the netplan configuration
func (np *Netplan) Apply() error {
	return fmt.Errorf("NOT Implemented")
}

// Store stores the updated network settings
func (np *Netplan) Store(ethernet Ethernet) error {
	if np.deviceNetplan.Network.Ethernets == nil {
		np.deviceNetplan.Network.Ethernets = map[string]Ethernet{ethernet.Name: ethernet}
	} else {
		np.deviceNetplan.Network.Ethernets[ethernet.Name] = ethernet
	}

	// Serialize the data to YAML
	data, err := yaml.Marshal(np.deviceNetplan)
	if err != nil {
		return nil
	}

	fmt.Println(string(data))
	return nil
}
