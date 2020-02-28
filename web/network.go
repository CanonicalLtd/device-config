// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"encoding/json"
	"fmt"
	"github.com/CanonicalLtd/configurator/service"
	"io"
	"net/http"
	"strings"
)

// InterfaceData defines the configuration of an interface
type InterfaceConfig struct {
	Use         bool     `json:"use"`
	Method      string   `json:"method"`
	Interface   string   `json:"interface"`
	NameServers []string `json:"nameServers"`
	Address     string   `json:"address"`
	Mask        string   `json:"mask"`
	Gateway     string   `json:"gateway"`
}

// Network is the API to get the network interface config
func (srv Web) Network(w http.ResponseWriter, r *http.Request) {
	// Get the current netplan settings
	netYAML := srv.Netplan.Current()

	// Get the hardware interfaces
	hardware, err := service.Interfaces()
	if err != nil {
		formatStandardResponse("interfaces", err.Error(), w)
		return
	}

	// Decode the configuration for each hardware interface
	interfaces := []InterfaceConfig{}
	for _, iface := range hardware {
		cfg := InterfaceConfig{Interface: iface.Name, Use: false}

		// Get the current interface config
		eth, ok := netYAML.Network.Ethernets[iface.Name]
		if !ok {
			// The interface is not configured
			interfaces = append(interfaces, cfg)
			continue
		}

		srv.decodeNetplanInterface(&cfg, eth)

		interfaces = append(interfaces, cfg)
	}

	// Create the JSON response
	formatNetworkResponse(interfaces, w)
}

func (srv Web) decodeNetplanInterface(cfg *InterfaceConfig, eth service.Ethernet) {
	// Parse the config
	cfg.Use = true
	cfg.Gateway = eth.Gateway4
	if len(eth.DHCP4) > 0 {
		cfg.Method = "dhcp"
	} else {
		cfg.Method = "manual"
	}
	if eth.NameServers != nil {
		cfg.NameServers = eth.NameServers["addresses"]
	}
	if eth.Addresses != nil {
		addressPlusMask := strings.Split(eth.Addresses[0], "/")
		cfg.Address = addressPlusMask[0]
		if len(addressPlusMask) > 1 {
			cfg.Mask = addressPlusMask[1]
		}
	}
}

func interfaces() []string {
	ifaces := []string{}
	interfaces, err := service.Interfaces()
	if err == nil {
		for _, iface := range interfaces {
			ifaces = append(ifaces, iface.Name)
		}
	}
	return ifaces
}

// NetworkInterface is the API to store the network interface configuration
func (srv Web) NetworkInterface(w http.ResponseWriter, r *http.Request) {
	req := srv.decodeNetworkInterface(w, r)
	if req == nil {
		return
	}

	eth := srv.encodeNetplanInterface(req)

	// Store the interface config
	if err := srv.Netplan.Store(eth); err != nil {
		formatStandardResponse("interface-store", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}

func (srv Web) encodeNetplanInterface(req *InterfaceConfig) service.Ethernet {
	// Encode the interface format into the netplan format
	eth := service.Ethernet{}
	eth.Name = req.Interface
	if req.Method == "dhcp" {
		eth.DHCP4 = "true"
	} else {
		eth.DHCP4 = ""
		eth.NameServers = map[string][]string{"addresses": req.NameServers}
		eth.Gateway4 = req.Gateway

		addr := req.Address
		if len(req.Mask) > 0 {
			addr = fmt.Sprintf("%s/%s", req.Address, req.Mask)
		}
		eth.Addresses = []string{addr}
	}
	return eth
}

func (srv Web) decodeNetworkInterface(w http.ResponseWriter, r *http.Request) *InterfaceConfig {
	// Decode the JSON body
	req := InterfaceConfig{}
	err := json.NewDecoder(r.Body).Decode(&req)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("data", "No interface data supplied.", w)
		return nil
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("decode-json", err.Error(), w)
		return nil
	}
	return &req
}
