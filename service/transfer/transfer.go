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

package transfer

import (
	"github.com/CanonicalLtd/device-config/service/dbus"
	"github.com/CanonicalLtd/device-config/service/snapd"
)

// Config holds the configuration data for export/import
type Config struct {
	NTP        bool   `yaml:"ntp"`
	Timezone   string `yaml:"timezone,omitempty"`
	ProxyHTTP  string `yaml:"http,omitempty"`
	ProxyHTTPS string `yaml:"https,omitempty"`
	ProxyFTP   string `yaml:"ftp,omitempty"`
}

// Service is the interface for export and import of the configuration
type Service interface {
	Export() (*Config, error)
	Import(cfg Config) error
}

// Transfer is the implementation of the transfer config service
type Transfer struct {
	DBusSrv dbus.Service
	SnapSrv snapd.Client
}

// NewTransfer creates a new transfer service
func NewTransfer(dbusSrv dbus.Service, snapSrv snapd.Client) *Transfer {
	return &Transfer{
		DBusSrv: dbusSrv,
		SnapSrv: snapSrv,
	}
}

// Export generates an export representation of the configuration
func (x *Transfer) Export() (*Config, error) {
	cfg := &Config{}

	// Get the time settings
	t := x.DBusSrv.TimeDateStatus()
	cfg.NTP = t.NTP
	cfg.Timezone = t.Timezone

	// Get the proxy settings
	system, err := x.SnapSrv.Conf("system")
	if err != nil {
		return cfg, err
	}
	if system["proxy"] == nil {
		return cfg, err
	}
	proxy, ok := system["proxy"].(map[string]interface{})
	if !ok {
		return cfg, err
	}
	if proxy["http"] != nil {
		cfg.ProxyHTTP = proxy["http"].(string)
	}
	if proxy["https"] != nil {
		cfg.ProxyHTTPS = proxy["https"].(string)
	}
	if proxy["ftp"] != nil {
		cfg.ProxyFTP = proxy["ftp"].(string)
	}

	return cfg, nil
}

// Import sets the configuration from the export format
func (x *Transfer) Import(cfg Config) error {
	// Set the time settings
	if cfg.Timezone != "" {
		_ = x.DBusSrv.SetTimezone(cfg.Timezone)
	}
	_ = x.DBusSrv.SetNTP(cfg.NTP)

	// Set the Proxy settings
	return x.SnapSrv.SetProxy(cfg.ProxyHTTP, cfg.ProxyHTTPS, cfg.ProxyFTP)
}
