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

// AppServices is the API to get the current application service status
func (srv Web) AppServices(w http.ResponseWriter, r *http.Request) {
	services, err := srv.Snapd.AppServices([]string{})
	if err != nil {
		formatStandardResponse("services", err.Error(), w)
		return
	}

	formatServicesResponse(services, w)
}
