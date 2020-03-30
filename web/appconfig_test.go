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

func TestWeb_AppConfig(t *testing.T) {
	tests := []struct {
		name       string
		wantStatus int
	}{
		{"valid", http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{})

			w := sendRequestWithAuth("GET", "/v1/config", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("AppServices() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
