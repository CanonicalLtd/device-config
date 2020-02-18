// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"fmt"
	"github.com/CanonicalLtd/configurator/config"
	"github.com/gorilla/mux"
	"net/http"
)

// Web implements the web service
type Web struct {
	Settings *config.Settings
	//BoardSrv service.Board
}

// NewWebService starts a new web service
func NewWebService(settings *config.Settings) *Web {
	return &Web{
		Settings: settings,
	}
}

// Start the web service
func (srv Web) Start() error {
	fmt.Printf("Starting service on port %s\n", srv.Settings.Port)
	return http.ListenAndServe(srv.Settings.Port, srv.Router())
}

// Router returns the application router
func (srv Web) Router() *mux.Router {
	// Start the web service router
	router := mux.NewRouter()

	//router.Handle("/v1/boards", Middleware(http.HandlerFunc(srv.BoardsList))).Methods("GET")
	//router.Handle("/v1/store/snaps/{snapName}", Middleware(http.HandlerFunc(srv.StoreSearchHandler))).Methods("POST")
	//router.Handle("/v1/build", Middleware(http.HandlerFunc(srv.Build))).Methods("POST")
	//router.Handle("/v1/build/fetch", Middleware(http.HandlerFunc(srv.GetLiveFSBuild))).Methods("POST")

	// Serve the static path
	//p := path.Join(srv.Settings.DocRoot, "/static/")
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Settings.DocRoot)))
	router.PathPrefix("/static/").Handler(fs)

	// Default path is the index page
	router.Handle("/", Middleware(http.HandlerFunc(srv.Index))).Methods("GET")
	router.NotFoundHandler = Middleware(http.HandlerFunc(srv.Index))

	return router
}
