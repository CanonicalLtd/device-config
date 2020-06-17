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

// SystemResources is the monitor of system resources
type SystemResources struct {
	CPU    float64 `json:"cpu"`
	Memory float64 `json:"memory"`
	Disk   float64 `json:"disk"`
}

// SystemResources monitors the system resources
func (srv Web) SystemResources(w http.ResponseWriter, r *http.Request) {
	cpu, err := srv.SystemSrv.CPU()
	if err != nil {
		formatStandardResponse("system", err.Error(), w)
		return
	}
	mem, err := srv.SystemSrv.Memory()
	if err != nil {
		formatStandardResponse("system", err.Error(), w)
		return
	}
	disk, err := srv.SystemSrv.Disk()
	if err != nil {
		formatStandardResponse("system", err.Error(), w)
		return
	}

	rec := SystemResources{
		CPU:    cpu,
		Memory: mem,
		Disk:   disk,
	}
	formatRecordResponse(rec, w)
}
