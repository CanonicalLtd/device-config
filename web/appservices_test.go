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

func TestWeb_AppServices(t *testing.T) {
	tests := []struct {
		name        string
		servicesErr bool
		wantStatus  int
	}{
		{"valid", false, http.StatusOK},
		{"invalid", true, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.ParseArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{false, false, tt.servicesErr}, &mockTime{})

			w := sendRequestWithAuth("GET", "/v1/services", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("AppServices() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
