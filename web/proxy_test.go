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
	"net/http"
	"testing"
)

func TestWeb_Proxy(t *testing.T) {
	tests := []struct {
		name       string
		confError  bool
		wantStatus int
	}{
		{"valid", false, http.StatusOK},
		{"invalid-interfaces", true, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{tt.confError, false, false}, &mockTime{})

			w := sendRequestWithAuth("GET", "/v1/proxy", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Proxy() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestWeb_ProxyUpdate(t *testing.T) {
	proxyHTTP := []byte(`{"http":"192.168.2.1:4000", "https":"192.168.2.1:4001"}`)
	proxyFTP := []byte(`{"ftp":"192.168.2.1:4002"}`)
	tests := []struct {
		name       string
		setconfErr bool
		data       []byte
		wantStatus int
	}{
		{"valid-http", false, proxyHTTP, http.StatusOK},
		{"valid-http-fail", true, proxyHTTP, http.StatusBadRequest},
		{"valid-ftp", false, proxyFTP, http.StatusOK},
		{"invalid-data", false, []byte(`\u1000`), http.StatusBadRequest},
		{"invalid-empty", false, nil, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{false, tt.setconfErr, false}, &mockTime{})

			w := sendRequestWithAuth("POST", "/v1/proxy", bytes.NewReader(tt.data), srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Proxy() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
