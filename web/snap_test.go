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

func TestWeb_SnapList(t *testing.T) {
	tests := []struct {
		name       string
		setErr     bool
		wantStatus int
	}{
		{"valid", false, http.StatusOK},
		{"invalid-interfaces", true, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{snapsErr: tt.setErr}, &mockTime{}, &mockSystem{})

				w := sendRequestWithAuth("GET", "/v1/snaps", nil, srv)
				if w.Code != tt.wantStatus {
					t.Errorf("SnapList() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
				}
			})
		})
	}
}

func TestWeb_SnapSet(t *testing.T) {
	validConf := []byte(`{"title":"Hello World"}`)
	invalidConf := []byte(`{'title':'Hello World''}`)

	tests := []struct {
		name       string
		setErr     bool
		data       []byte
		wantStatus int
	}{
		{"valid-conf", false, validConf, http.StatusOK},
		{"valid-conf-fail", true, validConf, http.StatusBadRequest},
		{"invalid-json", false, invalidConf, http.StatusBadRequest},
		{"invalid-data", false, []byte(`\u1000`), http.StatusBadRequest},
		{"invalid-empty", false, nil, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{setConfError: tt.setErr}, &mockTime{}, &mockSystem{})

				w := sendRequestWithAuth("PUT", "/v1/snaps/my-snap", bytes.NewReader(tt.data), srv)
				if w.Code != tt.wantStatus {
					t.Errorf("SnapSet() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
				}
			})
		})
	}
}
