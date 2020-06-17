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

package snapd

import (
	"encoding/json"
	"github.com/snapcore/snapd/client"
	"sync"
)

// Client is a client of the snapd REST API
type Client interface {
	AppServices(names []string) ([]*client.AppInfo, error)
	Conf(name string) (map[string]interface{}, error)
	List(names []string, opts *client.ListOptions) ([]Snap, error)
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

// AppServices requests the status of the application services
func (a *ClientAdapter) AppServices(names []string) ([]*client.AppInfo, error) {
	return a.snapdClient.Apps(names, client.AppOptions{Service: true})
}

// List returns the list of all snaps installed on the system
// with names in the given list; if the list is empty, all snaps.
func (a *ClientAdapter) List(names []string, opts *client.ListOptions) ([]Snap, error) {
	snaps, err := a.snapdClient.List(names, opts)
	if err != nil {
		return nil, err
	}

	ss := []Snap{}
	for _, s := range snaps {
		// Get the config for the snap (ignore errors)
		var conf string
		c, err := a.Conf(s.Name)
		if err == nil {
			resp, err := serializeResponse(c)
			if err == nil {
				conf = string(resp)
			}
		}

		ss = append(ss, Snap{
			Name:          s.Name,
			Title:         s.Title,
			Summary:       s.Summary,
			Description:   s.Description,
			InstalledSize: s.InstalledSize,
			InstalledDate: s.InstallDate,
			Status:        s.Status,
			Channel:       s.Channel,
			Confinement:   s.Confinement,
			Version:       s.Version,
			Revision:      s.Revision.N,
			Devmode:       s.DevMode,
			Config:        conf,
		})
	}
	return ss, nil
}

func serializeResponse(resp interface{}) ([]byte, error) {
	return json.Marshal(resp)
}
