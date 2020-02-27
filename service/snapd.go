/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

package service

import (
	"github.com/snapcore/snapd/client"
	"sync"
)

// SnapdClient is a client of the snapd REST API
type SnapdClient interface {
	Conf(name string) (map[string]interface{}, error)
	SetConf(name string, patch map[string]interface{}) (string, error)
}

// ClientAdapter adapts our expectations to the snapd client API.
type ClientAdapter struct {
	snapdClient *client.Client
}

var clientOnce sync.Once
var clientInstance *ClientAdapter

// NewClientAdapter creates a new ClientAdapter as a singleton
func NewClientAdapter() *ClientAdapter {
	clientOnce.Do(func() {
		clientInstance = &ClientAdapter{
			snapdClient: client.New(nil),
		}
	})

	return clientInstance
}

// Conf gets the snap's current configuration
func (a *ClientAdapter) Conf(name string) (map[string]interface{}, error) {
	return a.snapdClient.Conf(name, []string{})
}

// SetConf requests a snap to apply the provided patch to the configuration
func (a *ClientAdapter) SetConf(name string, patch map[string]interface{}) (string, error) {
	return a.snapdClient.SetConf(name, patch)
}
