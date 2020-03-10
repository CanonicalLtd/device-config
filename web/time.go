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
	"io"
	"net/http"
)

// TimeConfig allows update of the time config
type TimeConfig struct {
	NTP      bool   `json:"ntp"`
	Time     string `json:"time"`
	Timezone string `json:"timezone"`
}

// Time is the API for fetching the time config
func (srv Web) Time(w http.ResponseWriter, r *http.Request) {
	t := srv.TimeSrv.Current()

	formatTimeResponse(t, w)
}

// TimeConfig is the API for configuring the time
func (srv Web) TimeConfig(w http.ResponseWriter, r *http.Request) {
	t := srv.decodeTimeConfig(w, r)
	if t == nil {
		return
	}

	if err := srv.TimeSrv.Apply(t.NTP, t.Timezone, t.Time); err != nil {
		formatStandardResponse("time-config", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}

func (srv Web) decodeTimeConfig(w http.ResponseWriter, r *http.Request) *TimeConfig {
	// Decode the JSON body
	req := TimeConfig{}
	err := json.NewDecoder(r.Body).Decode(&req)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("data", "No time data supplied.", w)
		return nil
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("decode-json", err.Error(), w)
		return nil
	}
	return &req
}
