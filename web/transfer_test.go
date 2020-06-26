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

func TestWeb_TransferImport(t *testing.T) {
	cfg := []byte(`bnRwOiB0cnVlCnRpbWV6b25lOiBFdXJvcGUvTG9uZG9uCmh0dHBzOiAxOTIuMTY4LjIuMQo=`)
	cfgBad := []byte(`4YCF`)
	tests := []struct {
		name       string
		withErr    bool
		data       []byte
		wantStatus int
	}{
		{"valid-http", false, cfg, http.StatusOK},
		{"valid-bad-data", false, cfgBad, http.StatusBadRequest},
		{"invalid-data", false, []byte(`\u1000`), http.StatusBadRequest},
		{"invalid-empty", false, nil, http.StatusBadRequest},
		{"invalid-conf", true, cfg, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{}, &mockTransfer{withErr: tt.withErr})

			w := sendRequestWithAuth("POST", "/v1/transfer/import", bytes.NewReader(tt.data), srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Export() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}

func TestWeb_TransferExport(t *testing.T) {
	tests := []struct {
		name       string
		withErr    bool
		wantStatus int
	}{
		{"valid", false, http.StatusOK},
		{"with-error", true, http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := NewWebService(config.DefaultArgs(), &mockAuth{}, &mockNetplan{}, &mockSnapd{}, &mockTime{}, &mockSystem{}, &mockTransfer{withErr: tt.withErr})

			w := sendRequestWithAuth("GET", "/v1/transfer/export", nil, srv)
			if w.Code != tt.wantStatus {
				t.Errorf("Import() expected HTTP status '%d', got: %v", tt.wantStatus, w.Code)
			}
		})
	}
}
