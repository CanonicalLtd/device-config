// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"fmt"
	"github.com/CanonicalLtd/configurator/service"
	"log"
	"net/http"
	"strings"
)

type networkData struct {
	Method     string
	Interface  string
	Interfaces []string
	DNS        string
	Address    string
	Mask       string
	Gateway    string
	Common     commonData
}

// Network is the web page for configuring the network and proxy
func (srv Web) Network(w http.ResponseWriter, r *http.Request) {
	data := networkData{Method: "", Interfaces: interfaces(), Common: commonData{Username: getUsername(r)}}

	switch r.Method {
	case http.MethodPost:
		// Validate the settings
		data = networkData{
			Method:    r.FormValue("method"),
			Interface: r.FormValue("interface"),
			DNS:       r.FormValue("dns"),
			Address:   r.FormValue("address"),
			Mask:      r.FormValue("netmask"),
			Gateway:   r.FormValue("gateway"),
		}
		ethernet, err := srv.formToNetplan(data)
		if err != nil {
			data.Common.Error = err.Error()

			fmt.Println(data)

			//srv.netplanToForm(&data)
			srv.networkTemplate(w, data)
			return
		}

		// Store the settings
		srv.Netplan.Store(ethernet)
	}

	// Set up the form from the current netplan config
	srv.netplanToForm(&data)

	// Display the web form
	srv.networkTemplate(w, data)
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

func (srv Web) networkTemplate(w http.ResponseWriter, data networkData) {
	// Parse the templates
	t, err := srv.templates("network.html")
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (srv Web) formToNetplan(data networkData) (service.Ethernet, error) {
	if data.Method == "dhcp" {
		return service.Ethernet{
			Name:  data.Interface,
			DHCP4: "yes",
		}, nil
	}

	// Check the IP addresses
	if !service.ValidateIP(data.Address) {
		return service.Ethernet{}, fmt.Errorf("'address' is not a valid IP")
	}

	// Convert the comma-separated list into a slice and validate
	nameServers := strings.Split(data.DNS, ",")
	for _, ns := range nameServers {
		if !service.ValidateIP(ns) {
			return service.Ethernet{}, fmt.Errorf("'DNS' includes an valid IP")
		}
	}

	address := data.Address
	if len(data.Mask) > 0 {
		address = fmt.Sprintf("%s/%s", data.Address, data.Mask)
	}

	eth := service.Ethernet{
		Name:        data.Interface,
		Addresses:   []string{address},
		NameServers: map[string][]string{"addresses": nameServers},
		Gateway4:    data.Gateway,
	}
	return eth, nil
}

func (srv Web) netplanToForm(data *networkData) {
	// Get the current settings
	netYAML := srv.Netplan.Current()

	if len(netYAML.Network.Ethernets) == 0 {
		// Just use the current settings
		return
	}

	for k, eth := range netYAML.Network.Ethernets {
		data.Interface = k
		data.Gateway = eth.Gateway4
		if len(eth.DHCP4) > 0 {
			data.Method = "dhcp"
		} else {
			data.Method = "manual"
		}
		if eth.NameServers != nil {
			data.DNS = strings.Join(eth.NameServers["addresses"], ",")
		}
		if eth.Addresses != nil {
			addressPlusMask := strings.Split(eth.Addresses[0], "/")
			data.Address = addressPlusMask[0]
			if len(addressPlusMask) > 0 {
				data.Mask = addressPlusMask[1]
			}
		}

		break
	}
}
