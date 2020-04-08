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

import "net/http"

// AppConfig holds the config for application
type AppConfig struct {
	SnapControl    bool     `json:"snapControl"`
	HideInterfaces []string `json:"hideInterfaces"`
}

// AppConfig is the API to get the application config
func (srv Web) AppConfig(w http.ResponseWriter, r *http.Request) {
	// Get the current settings
	cfg := AppConfig{SnapControl: srv.Settings.SnapControl, HideInterfaces: srv.Settings.HideInterfaces}

	// Return the response - snapd returns in AppConfig format
	formatAppConfigResponse(cfg, w)
}
