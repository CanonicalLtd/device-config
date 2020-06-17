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

func TestWeb_Time(t *testing.T) {
	tests := []struct {
		name       string
		wantStatus int
	}{
		{"valid", http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{})

			w := sendRequestWithAuth("GET", "/v1/time", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Time() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestWeb_TimeConfig(t *testing.T) {
	timeSync := []byte(`{"timezone":"America/Hawaii", "ntp":true}`)
	timeZoneInvalid := []byte(`{"timezone":"invalid", "ntp":true}`)
	tests := []struct {
		name       string
		data       []byte
		wantStatus int
	}{
		{"valid-time-sync", timeSync, http.StatusOK},
		{"valid-time-invalid", timeZoneInvalid, http.StatusBadRequest},
		{"invalid-data", []byte(`\u1000`), http.StatusBadRequest},
		{"invalid-empty", nil, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{})

			w := sendRequestWithAuth("POST", "/v1/time", bytes.NewReader(tt.data), srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Time() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
