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

package memory

import (
	"github.com/CanonicalLtd/device-config/datastore"
	"testing"
	"time"
)

func TestStore_CreateSession(t *testing.T) {
	tests := []struct {
		name    string
		user    datastore.Session
		want    int64
		wantErr bool
	}{
		{"valid", datastore.Session{Username: "jsmith", SessionID: "abc", Expires: time.Now().Add(365 * 24 * time.Hour)}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mem := NewStore()
			got, err := mem.CreateSession(tt.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateSession() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_GetSession(t *testing.T) {
	// Set up the user
	mem := NewStore()
	id, err := mem.CreateSession(datastore.Session{Username: "jsmith", SessionID: "abc", Expires: time.Now().Add(365 * 24 * time.Hour)})
	if err != nil {
		t.Errorf("GetSession() error creating user = %v, wantErr %v", err, nil)
	}
	if id != 1 {
		t.Errorf("GetSession() got = %v, want %v", id, 1)
	}
	_, err = mem.CreateSession(datastore.Session{Username: "expired", SessionID: "abc", Expires: time.Now()})
	if err != nil {
		t.Errorf("GetSession() error creating user = %v, wantErr %v", err, nil)
	}

	type args struct {
		username  string
		sessionID string
	}
	tests := []struct {
		name    string
		args    args
		wantID  int64
		wantErr bool
	}{
		{"expired", args{"expired", "abc"}, 0, true},
		{"valid", args{"jsmith", "abc"}, id, false},
		{"invalid-user", args{"invalid", "abc"}, 0, true},
		{"invalid-session", args{"jsmith", "invalid"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mem.GetSession(tt.args.username, tt.args.sessionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.ID != tt.wantID {
				t.Errorf("GetSession() got = %v, want %v", got.ID, tt.wantID)
			}
		})
	}
}
