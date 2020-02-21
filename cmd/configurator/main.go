// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package main

import (
	"github.com/CanonicalLtd/configurator/config"
	"github.com/CanonicalLtd/configurator/datastore/memory"
	"github.com/CanonicalLtd/configurator/service"
	"github.com/CanonicalLtd/configurator/web"
	"log"
)

func main() {
	// Parse the command-line arguments
	settings := config.ParseArgs()

	// Set up the dependency chain
	memorySrv := memory.NewStore()
	authSrv := service.NewAuthService(memorySrv)
	srv := web.NewWebService(settings, authSrv)

	// Start the web service
	log.Fatal(srv.Start())
}
