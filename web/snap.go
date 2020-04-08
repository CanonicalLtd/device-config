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
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// SnapList list the installed snaps
func (srv Web) SnapList(w http.ResponseWriter, r *http.Request) {
	snaps, err := srv.Snapd.List([]string{}, nil)
	if err != nil {
		formatStandardResponse("snaps", err.Error(), w)
		return
	}

	formatRecordsResponse(snaps, w)
}

// SnapSet updates the config of a snap
func (srv Web) SnapSet(w http.ResponseWriter, r *http.Request) {
	data := srv.decodeSnapConfig(w, r)
	if data == nil {
		return
	}

	vars := mux.Vars(r)
	_, err := srv.Snapd.SetConf(vars["snap"], data)
	if err != nil {
		formatStandardResponse("snap-set", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}

func (srv Web) decodeSnapConfig(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// Decode the JSON body
	req := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&req)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("data", "No proxy data supplied.", w)
		return nil
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("decode-json", err.Error(), w)
		return nil
	}
	return req
}
