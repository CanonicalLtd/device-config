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

import "time"

// Snap holds the details of snap on a device
type Snap struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	Summary       string    `json:"summary"`
	Description   string    `json:"description"`
	InstalledSize int64     `json:"installedSize"`
	InstalledDate time.Time `json:"installedDate"`
	Status        string    `json:"status"`
	Channel       string    `json:"channel"`
	Confinement   string    `json:"confinement"`
	Version       string    `json:"version"`
	Revision      int       `json:"revision"`
	Devmode       bool      `json:"devmode"`
	Config        string    `json:"config"`
}
