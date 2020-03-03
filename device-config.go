// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package main

import (
	"github.com/CanonicalLtd/device-config/config"
	"github.com/CanonicalLtd/device-config/datastore/memory"
	"github.com/CanonicalLtd/device-config/service"
	"github.com/CanonicalLtd/device-config/web"
	"log"
)

func main() {
	// Parse the command-line arguments
	settings := config.ParseArgs()

	// Set up the dependency chain
	memorySrv := memory.NewStore()
	authSrv := service.NewAuthService(memorySrv)
	snapdClient := service.NewClientAdapter()
	dBus, err := service.NewDBus()
	if err != nil {
		log.Fatal(err)
	}
	netplanSrv := service.NewNetplan(dBus)
	timeSrv := service.NewTime(dBus)
	srv := web.NewWebService(settings, authSrv, netplanSrv, snapdClient, timeSrv)

	// Start the web service
	log.Fatal(srv.Start())
}
