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
	"strings"
	"time"
)

// loginData is the data for a login response
type loginData struct {
	MacAddress string `json:"macAddress"`
}

// Login is the login web page
func (srv Web) Login(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON body
	data := loginData{}
	err := json.NewDecoder(r.Body).Decode(&data)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("login-data", "No login data supplied", w)
		return
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("decode-json", err.Error(), w)
		return
	}

	// Allow using `-` instead of `:`
	data.MacAddress = strings.ReplaceAll(data.MacAddress, "-", ":")

	// Validate the request
	expires := time.Now().Add(24 * time.Hour)
	session, err := srv.Auth.CreateSession(data.MacAddress, expires)
	if err != nil {
		formatStandardResponse("login", err.Error(), w)
		return
	}

	// Add the session cookie
	setCookies(w, session.Username, session.SessionID, expires)
	formatLoginResponse(session.Username, session.SessionID, w)
}

// Logout is the logout web page to remove the session cookie
func (srv Web) Logout(w http.ResponseWriter, r *http.Request) {
	// Update the cookies
	expiration := time.Unix(0, 0)
	cookie1 := http.Cookie{Name: "username", Value: "", Expires: expiration}
	cookie2 := http.Cookie{Name: "sessionID", Value: "", Expires: expiration}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func setCookies(w http.ResponseWriter, username, sessionID string, expiration time.Time) {
	cookie1 := http.Cookie{Name: "username", Value: username, Expires: expiration}
	cookie2 := http.Cookie{Name: "sessionID", Value: sessionID, Expires: expiration}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
}
