// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"github.com/CanonicalLtd/device-config/datastore"
	"log"
	"net/http"
	"time"
)

// Logger Handle logging for the web service
func Logger(start time.Time, r *http.Request) {
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

// Middleware to pre-process web service requests
func Middleware(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		Logger(start, r)

		inner.ServeHTTP(w, r)
	})
}

// AuthCheck checks that we have an active session cookie
func (srv Web) AuthCheck(r *http.Request) (*datastore.Session, error) {
	// Get the session cookies
	username, err := r.Cookie("username")
	if err != nil {
		return nil, err
	}
	sessionID, err := r.Cookie("sessionID")
	if err != nil {
		return nil, err
	}

	// Check that the session details are valid
	return srv.Auth.ValidateSession(username.Value, sessionID.Value)
}

// MiddlewareWithAuth handles authentication and redirects to the login page
func (srv Web) MiddlewareWithAuth(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		Logger(start, r)

		// Check that we have the session cookies and that they are valid
		_, err := srv.AuthCheck(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// TODO: refresh the session cookie

		inner.ServeHTTP(w, r)
	})
}
