// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"log"
	"net/http"
	"time"
)

// loginData is the data for a login response
type loginData struct {
	Common commonData
}

// Login is the login web page
func (srv Web) Login(w http.ResponseWriter, r *http.Request) {
	data := loginData{commonData{Username: getUsername(r)}}

	// Handle a submitted form
	if r.Method == http.MethodPost {
		// Validate the form
		macAddr := r.FormValue("macaddress")
		session, err := srv.Auth.CreateSession(macAddr)
		if err != nil {
			data.Common.Error = err.Error()
		} else {
			// Add the session cookie
			setCookies(w, session.Username, session.SessionID)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	// Parse the templates
	t, err := srv.templates("login.html")
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
