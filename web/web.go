// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"fmt"
	"github.com/CanonicalLtd/configurator/config"
	"github.com/CanonicalLtd/configurator/service"
	"github.com/gorilla/mux"
	"net/http"
)

// Web implements the web service
type Web struct {
	Settings *config.Settings
	Auth     service.AuthService
	Netplan  service.NetplanService
}

// NewWebService starts a new web service
func NewWebService(settings *config.Settings, auth service.AuthService, netplan service.NetplanService) *Web {
	return &Web{
		Settings: settings,
		Auth:     auth,
		Netplan:  netplan,
	}
}

// Start the web service
func (srv Web) Start() error {
	listenOn := fmt.Sprintf("%s:%s", srv.Settings.NetworkInterface, srv.Settings.Port)
	fmt.Printf("Starting service on port %s\n", listenOn)
	return http.ListenAndServe(listenOn, srv.Router())
}

// Router returns the application router
func (srv Web) Router() *mux.Router {
	// Start the web service router
	router := mux.NewRouter()

	router.Handle("/login", Middleware(http.HandlerFunc(srv.Login))).Methods("GET", "POST")
	router.Handle("/logout", Middleware(http.HandlerFunc(srv.Logout))).Methods("GET")
	router.Handle("/network", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Network))).Methods("GET", "POST")
	router.Handle("/time", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Time))).Methods("GET")

	// Serve the static path
	//p := path.Join(srv.Settings.DocRoot, "/static/")
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Settings.DocRoot)))
	router.PathPrefix("/static/").Handler(fs)

	// Default path is the index page
	router.Handle("/", Middleware(http.HandlerFunc(srv.Index))).Methods("GET")
	router.NotFoundHandler = Middleware(http.HandlerFunc(srv.Index))

	return router
}
