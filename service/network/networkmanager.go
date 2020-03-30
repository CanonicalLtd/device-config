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
	"log"
)

// NMService is the interface for the netplan service
type NMService interface {
	Apply() error
	Current() *NetplanYAML
	Store(ethernet Ethernet) error
}

// NetManager implements actions for managing network manager
type NetManager struct {
	deviceNetplan *NetplanYAML
	interfaces    map[string]string
	dBus          dbus.Service
}

// NewNetworkManager creates a network manager service
func NewNetworkManager(dBus dbus.Service) *NetManager {
	fmt.Println("Using network manager for network configuration")
	deviceNetplan := &NetplanYAML{Network: Network{Version: 2, Renderer: "NetworkManager", Ethernets: map[string]Ethernet{}}}

	// Write the netplan file for network-manager
	_ = serializeNetplan(deviceNetplan)

	// Get the devices from dbus
	devices, err := dBus.NMDevices()
	if err != nil {
		devices = map[string]string{}
	}
	return &NetManager{dBus: dBus, interfaces: devices, deviceNetplan: deviceNetplan}
}

// Current returns the current network manager settings
func (np *NetManager) Current() *NetplanYAML {
	// Get the config details for the active connections
	interfaces, err := Interfaces()
	if err != nil {
		fmt.Println("Error fetching current settings:", err)
		return np.deviceNetplan
	}

	// Map the NM details to our internal format
	for _, iface := range interfaces {
		p := np.interfaces[iface.Name]
		ifaceConfig := np.dBus.NMInterfaceConfig(p)

		dhcp4 := ""
		if ifaceConfig.DHCP4 {
			dhcp4 = "true"
		}
		addresses := []string{}
		for _, a := range ifaceConfig.AddressData {
			addresses = append(addresses, fmt.Sprintf("%s/%d", a.Address, a.Prefix))
		}

		eth := Ethernet{
			Use:         ifaceConfig.State >= 100,
			Name:        iface.Name,
			DHCP4:       dhcp4,
			Addresses:   addresses,
			NameServers: map[string][]string{"addresses": ifaceConfig.NameServers},
			Gateway4:    ifaceConfig.Gateway,
		}
		np.deviceNetplan.Network.Ethernets[iface.Name] = eth
	}
	log.Println("---", np.deviceNetplan)
	return np.deviceNetplan
}

// Apply applies the network manager configuration using dbus
func (np *NetManager) Apply() error {
	// This is not needed for network-manager connections
	return nil
}

// Store stores the updated network settings
func (np *NetManager) Store(ethernet Ethernet) error {
	var eth dbus.NMDeviceSettings

	//  Get the dbus path for the interface
	p := np.interfaces[ethernet.Name]

	eth = dbus.NMDeviceSettings{
		DHCP4: ethernet.DHCP4 == "true",
	}
	if ethernet.DHCP4 == "true" {
		// Configure the device for DHCP
		_ = np.dBus.NMInterfaceConfigUpdate(p, eth)
		return nil
	}

	// Convert the format of the settings
	addressData := []dbus.NMDeviceAddress{}
	for _, a := range ethernet.Addresses {
		ip, mask, err := validateAddress(a)
		if err != nil {
			return err
		}

		addressData = append(addressData, dbus.NMDeviceAddress{Address: ip, Prefix: mask})
	}

	eth.AddressData = addressData
	eth.NameServers = ethernet.NameServers["addresses"]
	eth.Gateway = ethernet.Gateway4

	return np.dBus.NMInterfaceConfigUpdate(p, eth)
}
