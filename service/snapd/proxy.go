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

// SetProxy sets the proxy configuration
func (a *ClientAdapter) SetProxy(http, https, ftp string) error {
	// Set up the proxy settings
	settings := map[string]interface{}{}
	if len(http) > 0 {
		settings["http"] = http
	} else {
		settings["http"] = nil
	}
	if len(https) > 0 {
		settings["https"] = https
	} else {
		settings["https"] = nil
	}
	if len(ftp) > 0 {
		settings["ftp"] = ftp
	} else {
		settings["ftp"] = nil
	}

	cfg := map[string]interface{}{"proxy": settings}

	// Save the proxy settings
	_, err := a.SetConf("system", cfg)
	return err
}
