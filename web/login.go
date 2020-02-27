// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"encoding/json"
	"io"
	"net/http"
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

	// Validate the request
	session, err := srv.Auth.CreateSession(data.MacAddress)
	if err != nil {
		formatStandardResponse("login", err.Error(), w)
		return
	}

	// Add the session cookie
	setCookies(w, session.Username, session.SessionID)
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

func setCookies(w http.ResponseWriter, username, sessionID string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie1 := http.Cookie{Name: "username", Value: username, Expires: expiration}
	cookie2 := http.Cookie{Name: "sessionID", Value: sessionID, Expires: expiration}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
}
