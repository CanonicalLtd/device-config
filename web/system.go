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

import (
	"encoding/json"
	"github.com/CanonicalLtd/device-config/service"
	"io"
	"net/http"
	"strings"
)

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

// FactoryReset triggers a factory reset on the device
func (srv Web) FactoryReset(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body
	data := loginData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("reset-data", "No reset data supplied", w)
		return
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("decode-json", err.Error(), w)
		return
	}

	// Allow using `-` instead of `:`
	data.MacAddress = strings.ReplaceAll(data.MacAddress, "-", ":")

	// Check that the entered MAC address is valid
	if err := service.CheckMacAddress(data.MacAddress); err != nil {
		formatStandardResponse("reset", err.Error(), w)
		return
	}

	// Trigger the factory reset
	if err := srv.SystemSrv.FactoryReset(); err != nil {
		formatStandardResponse("reset", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}
