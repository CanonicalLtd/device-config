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
	"github.com/CanonicalLtd/device-config/config"
	"github.com/CanonicalLtd/device-config/service"
	"github.com/gorilla/mux"
	"net/http"
)

// Web implements the web service
type Web struct {
	Settings *config.Settings
	Auth     service.AuthService
	Netplan  service.NetplanService
	Snapd    service.SnapdClient
	TimeSrv  service.TimeService
}

// NewWebService starts a new web service
func NewWebService(settings *config.Settings, auth service.AuthService, netplan service.NetplanService, snapd service.SnapdClient, t service.TimeService) *Web {
	return &Web{
		Settings: settings,
		Auth:     auth,
		Netplan:  netplan,
		Snapd:    snapd,
		TimeSrv:  t,
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

	router.Handle("/v1/config", Middleware(http.HandlerFunc(srv.AppConfig))).Methods("GET")
	router.Handle("/v1/login", Middleware(http.HandlerFunc(srv.Login))).Methods("POST")
	router.Handle("/v1/network", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Network))).Methods("GET")
	router.Handle("/v1/network", srv.MiddlewareWithAuth(http.HandlerFunc(srv.NetworkInterface))).Methods("POST")
	router.Handle("/v1/network/apply", srv.MiddlewareWithAuth(http.HandlerFunc(srv.NetworkApply))).Methods("POST")
	router.Handle("/v1/proxy", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Proxy))).Methods("GET")
	router.Handle("/v1/proxy", srv.MiddlewareWithAuth(http.HandlerFunc(srv.ProxyUpdate))).Methods("POST")
	router.Handle("/v1/time", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Time))).Methods("GET")
	router.Handle("/v1/time", srv.MiddlewareWithAuth(http.HandlerFunc(srv.TimeConfig))).Methods("POST")
	router.Handle("/v1/services", srv.MiddlewareWithAuth(http.HandlerFunc(srv.AppServices))).Methods("GET")
	router.Handle("/logout", Middleware(http.HandlerFunc(srv.Logout))).Methods("GET")

	// Serve the static path
	//p := path.Join(srv.Settings.DocRoot, "/static/")
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(srv.Settings.DocRoot)))
	router.PathPrefix("/static/").Handler(fs)

	// Default path is the index page
	router.Handle("/", Middleware(http.HandlerFunc(srv.Index))).Methods("GET")
	router.Handle("/login", Middleware(http.HandlerFunc(srv.Index))).Methods("GET")
	router.Handle("/network", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Index))).Methods("GET")
	router.Handle("/proxy", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Index))).Methods("GET")
	router.Handle("/time", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Index))).Methods("GET")
	router.Handle("/services", srv.MiddlewareWithAuth(http.HandlerFunc(srv.Index))).Methods("GET")
	router.NotFoundHandler = Middleware(http.HandlerFunc(srv.Index))

	return router
}
