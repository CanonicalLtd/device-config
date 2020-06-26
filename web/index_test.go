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
	"github.com/CanonicalLtd/device-config/config"
	"net/http"
	"testing"
)

func TestWeb_Index(t *testing.T) {
	tests := []struct {
		name       string
		template   string
		wantStatus int
	}{
		{"valid", config.DefaultIndexTemplate, http.StatusOK},
		{"invalid-path", "does-not-exist.html", http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.DefaultArgs()
			cfg.DocRoot = "../static"
			cfg.IndexTemplate = tt.template
			srv := NewWebService(cfg, &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{}, &mockTransfer{})

			w := sendRequest("GET", "/", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Web() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestWeb_NoAuth(t *testing.T) {
	// Check the redirect to the login when the user is not authenticated
	tests := []struct {
		name       string
		url        string
		wantStatus int
	}{
		{"network-api", "/v1/network", http.StatusSeeOther},
		{"network", "/network", http.StatusSeeOther},
		{"services-api", "/v1/services", http.StatusSeeOther},
		{"services", "/services", http.StatusSeeOther},
		{"proxy-api", "/v1/proxy", http.StatusSeeOther},
		{"proxy", "/proxy", http.StatusSeeOther},
		{"time-api", "/v1/time", http.StatusSeeOther},
		{"time", "/time", http.StatusSeeOther},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{}, &mockTransfer{})

			w := sendRequest("GET", tt.url, nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("NoAuth expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
