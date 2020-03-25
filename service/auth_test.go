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

package service

import (
	"fmt"
	"github.com/CanonicalLtd/device-config/datastore"
	"github.com/CanonicalLtd/device-config/datastore/memory"
	"github.com/CanonicalLtd/device-config/service/network"
	"testing"
	"time"
)

func mockInterfacesValid() ([]network.NetworkInterface, error) {
	return []network.NetworkInterface{{"eth0", "eth0"}}, nil
}
func mockInterfacesNone() ([]network.NetworkInterface, error) {
	return []network.NetworkInterface{}, fmt.Errorf("MOCK error")
}

func TestAuth_CreateSession(t *testing.T) {
	type args struct {
		token   string
		expires time.Time
	}
	tests := []struct {
		name       string
		args       args
		ifaceError bool
		wantErr    bool
	}{
		{"valid", args{"eth0", time.Now()}, false, false},
		{"valid-not-found", args{"not-found", time.Now()}, false, true},
		{"invalid-interfaces", args{"abc", time.Now()}, true, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the interfaces call
			if tt.ifaceError {
				network.Interfaces = mockInterfacesNone
			} else {
				network.Interfaces = mockInterfacesValid
			}

			mem := memory.NewStore()
			auth := NewAuthService(mem)
			got, err := auth.CreateSession(tt.args.token, tt.args.expires)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if len(got.Username) == 0 || len(got.SessionID) == 0 {
				t.Errorf("CreateSession() session, got %v", got)
				return
			}
		})
	}
}

func TestAuth_ValidateSession(t *testing.T) {
	// Set up the data store
	mem := memory.NewStore()
	mem.CreateSession(datastore.Session{
		Username:  "jsmith",
		SessionID: "abc123",
		Expires:   time.Now().Add(24 * time.Hour),
	})

	type args struct {
		username  string
		sessionID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{"jsmith", "abc123"}, false},
		{"invalid", args{"invalid", "abc123"}, true},
		{"invalid-session", args{"jsmith", "invalid"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			auth := NewAuthService(mem)
			got, err := auth.ValidateSession(tt.args.username, tt.args.sessionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}

			if got.Username != tt.args.username || got.SessionID != tt.args.sessionID {
				t.Errorf("ValidateSession() got = %v, want %v", got, tt.args)
			}
		})
	}
}
