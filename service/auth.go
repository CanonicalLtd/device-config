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
	"github.com/google/uuid"
	"strings"
)

// AuthService is the interface for the authentication service
type AuthService interface {
	ValidateSession(username, sessionID string) (*datastore.Session, error)
	CreateSession(token string) (*datastore.Session, error)
}

// Auth is the implementation of the authentication service
type Auth struct {
	DataStore datastore.DataStore
}

// NewAuthService creates a new authentication service
func NewAuthService(store datastore.DataStore) *Auth {
	return &Auth{
		DataStore: store,
	}
}

// CreateSession creates a new session, validating the token
func (auth *Auth) CreateSession(token string) (*datastore.Session, error) {
	// Check the token against the MAC addresses
	if err := auth.checkMacAddress(token); err != nil {
		return nil, err
	}

	// Create the user session
	user := datastore.Session{
		Username:  uuid.New().String(),
		SessionID: uuid.New().String(),
	}
	_, err := auth.DataStore.CreateSession(user)
	return &user, err
}

// ValidateSession checks that the session is valid
func (auth *Auth) ValidateSession(username, sessionID string) (*datastore.Session, error) {
	return auth.DataStore.GetSession(username, sessionID)
}

func (auth *Auth) checkMacAddress(token string) error {
	// Get the hardware interfaces
	interfaces, err := Interfaces()
	if err != nil {
		return err
	}

	// Check that we have a matching MAC address
	for _, n := range interfaces {
		if strings.EqualFold(n.MACAddress, token) {
			return nil
		}
	}

	return fmt.Errorf("could not find a matching MAC address")
}
