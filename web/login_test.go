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

func TestWeb_Login(t *testing.T) {
	loginOk := []byte(`{"macAddress":"valid-mac-address"}`)
	loginInvalid := []byte(`{"macAddress":"invalid"}`)
	tests := []struct {
		name       string
		data       []byte
		wantStatus int
	}{
		{"valid", loginOk, http.StatusOK},
		{"invalid-login", loginInvalid, http.StatusBadRequest},
		{"invalid-data", []byte(`\u1000`), http.StatusBadRequest},
		{"invalid-empty", nil, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{})

			w := sendRequest("POST", "/v1/login", bytes.NewReader(tt.data), srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Login() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
			if w.Code != http.StatusOK {
				return
			}
			if w.Header().Get("Set-Cookie") == "" {
				t.Errorf("Login() cookie error: username '%v', session ID: '%v'", w.Header().Get("username"), w.Header().Get("sessionID"))
			}
		})
	}
}

func TestWeb_Logout(t *testing.T) {
	srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{})

	w := sendRequest("GET", "/logout", nil, srv)
	if w.Code != http.StatusSeeOther {
		t.Errorf("Logout() expected HTTP status '%v', got: %v", http.StatusSeeOther, w.Code)
	}
}
