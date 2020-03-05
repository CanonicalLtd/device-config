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
	"github.com/snapcore/snapd/client"
	"sync"
)

// SnapdClient is a client of the snapd REST API
type SnapdClient interface {
	Conf(name string) (map[string]interface{}, error)
	SetConf(name string, patch map[string]interface{}) (string, error)
	AppServices(names []string) ([]*client.AppInfo, error)
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

// AppServices requests the status of the application services
func (a *ClientAdapter) AppServices(names []string) ([]*client.AppInfo, error) {
	return a.snapdClient.Apps(names, client.AppOptions{Service: true})
}
