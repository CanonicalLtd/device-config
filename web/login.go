// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// loginData is the data for a login response
type loginData struct {
	Error string
}

// Login is the login web page
func (srv Web) Login(w http.ResponseWriter, r *http.Request) {
	data := loginData{}

	// Handle a submitted form
	if r.Method == http.MethodPost {
		// Validate the form
		macAddr := r.FormValue("macaddress")
		fmt.Println("------", macAddr)
		session, err := srv.Auth.CreateSession(macAddr)
		if err != nil {
			fmt.Println("---", err.Error())
			data.Error = err.Error()
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

func setCookies(w http.ResponseWriter, username, sessionID string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie1 := http.Cookie{Name: "username", Value: username, Expires: expiration}
	cookie2 := http.Cookie{Name: "sessionID", Value: sessionID, Expires: expiration}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
}
