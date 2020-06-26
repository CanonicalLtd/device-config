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
	"testing"
)

func TestWeb_SystemResources(t *testing.T) {
	tests := []struct {
		name       string
		cpuErr     bool
		memErr     bool
		diskErr    bool
		wantStatus int
	}{
		{"valid", false, false, false, http.StatusOK},
		{"cpu-error", true, false, false, http.StatusBadRequest},
		{"mem-error", false, true, false, http.StatusBadRequest},
		{"disk-error", false, false, true, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{tt.cpuErr, tt.memErr, tt.diskErr, false}, &mockTransfer{})

			w := sendRequestWithAuth("GET", "/v1/system", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Proxy() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestWeb_FactoryReset(t *testing.T) {
	resetOk := []byte(`{"macAddress":"eth1-mac-address"}`)
	resetInvalid := []byte(`{"macAddress":"invalid"}`)
	tests := []struct {
		name       string
		data       []byte
		resetErr   bool
		wantStatus int
	}{
		{"valid", resetOk, false, http.StatusOK},
		{"valid-with-error", resetOk, true, http.StatusBadRequest},
		{"invalid-reset", resetInvalid, false, http.StatusBadRequest},
		{"invalid-data", []byte(`\u1000`), false, http.StatusBadRequest},
		{"invalid-empty", nil, false, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the retrieval of network interfaces
			network.Interfaces = mockInterfacesValid

			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{resetErr: tt.resetErr}, &mockTransfer{})

			w := sendRequest("POST", "/v1/factory-reset", bytes.NewReader(tt.data), srv)
			if w.Code != tt.wantStatus {
				t.Errorf("FactoryReset() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
