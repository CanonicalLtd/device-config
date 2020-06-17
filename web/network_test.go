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
	"bytes"
	"github.com/CanonicalLtd/device-config/config"
	"github.com/CanonicalLtd/device-config/service/network"
	"net/http"
	"strings"
	"testing"
)

func TestWeb_Network(t *testing.T) {
	tests := []struct {
		name       string
		ifaceErr   bool
		wantStatus int
	}{
		{"valid", false, http.StatusOK},
		{"invalid-interfaces", true, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the retrieval of network interfaces
			if tt.ifaceErr {
				network.Interfaces = mockInterfacesNone
			} else {
				network.Interfaces = mockInterfacesValid
			}

			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{})

			w := sendRequestWithAuth("GET", "/v1/network", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Network() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestWeb_NetworkInterface(t *testing.T) {
	netDHCP := []byte(`{"use":true, "method":"dhcp", "interface":"eth0"}`)
	netManual := []byte(`{"use":true, "method":"manual", "interface":"eth0", "nameServers":["192.168.2.1","8.8.8.8"], "address":"192.168.1.100", "mask":"192.168.1.255","gateway":"192.168.1.1"}`)
	netBadIface := []byte(`{"use":true, "method":"dhcp", "interface":"invalid"}`)

	tests := []struct {
		name       string
		ifaceErr   bool
		data       []byte
		wantStatus int
	}{
		{"valid-dhcp", false, netDHCP, http.StatusOK},
		{"valid-manual", false, netManual, http.StatusOK},
		{"invalid-data", false, []byte(`\u1000`), http.StatusBadRequest},
		{"invalid-empty", false, nil, http.StatusBadRequest},
		{"invalid-interface", false, netBadIface, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the retrieval of network interfaces
			if tt.ifaceErr {
				network.Interfaces = mockInterfacesNone
			} else {
				network.Interfaces = mockInterfacesValid
			}

			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{})
			w := sendRequestWithAuth("POST", "/v1/network", bytes.NewReader(tt.data), srv)
			if w.Code != tt.wantStatus {
				t.Errorf("HardwareInterface() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestWeb_NetworkApply(t *testing.T) {
	tests := []struct {
		name       string
		applyError bool
		wantStatus int
	}{
		{"valid", false, http.StatusOK},
		{"valid", true, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{tt.applyError}, &mockSnapd{}, &mockTime{}, &mockSystem{})
			w := sendRequestWithAuth("POST", "/v1/network/apply", strings.NewReader(""), srv)
			if w.Code != tt.wantStatus {
				t.Errorf("NetworkApply() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
