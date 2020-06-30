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
	"log"
	"net/http"
)

type indexData struct {
	Title  string
	Common commonData
}

// Index is the front page of the web application
func (srv Web) Index(w http.ResponseWriter, r *http.Request) {
	data := indexData{Title: srv.Settings.Custom.Title, Common: commonData{Username: getUsername(r)}}

	t, err := srv.templates(srv.Settings.IndexTemplate)
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
