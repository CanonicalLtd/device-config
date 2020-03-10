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
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func sendRequest(method, url string, data io.Reader, srv *Web) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, data)

	srv.Router().ServeHTTP(w, r)
	return w
}

func sendRequestWithAuth(method, url string, data io.Reader, srv *Web) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest(method, url, data)

	// Set a valid cookie
	c1 := fmt.Sprintf("username=%s; ", "generated-user")
	c2 := fmt.Sprintf("sessionID=%s", "generated-session-id")
	r.Header.Set("Cookie", c1+c2)

	srv.Router().ServeHTTP(w, r)
	return w
}
