// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package main

import (
	"flag"
	"fmt"
	"github.com/CanonicalLtd/device-config/config"
	"github.com/CanonicalLtd/device-config/datastore/memory"
	"github.com/CanonicalLtd/device-config/service"
	"github.com/CanonicalLtd/device-config/service/dbus"
	"github.com/CanonicalLtd/device-config/service/network"
	"github.com/CanonicalLtd/device-config/web"
	"log"
	"os"
)

func main() {
	// Get the application parameters
	settings := config.ReadParameters()
	configure(settings)

	// Set up the dependency chain
	memorySrv := memory.NewStore()
	authSrv := service.NewAuthService(memorySrv)
	snapdClient := service.NewClientAdapter()
	dBus, err := dbus.NewDBus()
	if err != nil {
		log.Fatal(err)
	}
	netplanSrv := network.NewNetplan(dBus)
	timeSrv := service.NewTime(dBus)

	nm := network.NewNetworkManager(dBus)
	cfg := nm.Current()
	fmt.Println("---", *cfg)

	srv := web.NewWebService(settings, authSrv, netplanSrv, snapdClient, timeSrv)

	// Start the web service
	log.Fatal(srv.Start())
}

func configure(cfg *config.Settings) {
	var (
		configureOnly bool
		iface         string
		manageProxy   bool
	)
	flag.BoolVar(&configureOnly, "configure", false, "Configure the application and exit")
	flag.StringVar(&iface, "interface", config.DefaultInterface, "The default network interface for the service")
	flag.BoolVar(&manageProxy, "proxy", config.DefaultManageProxy, "Allow proxy configuration (needs the snapd-control interface)")
	flag.Parse()

	log.Printf("Device config: configure=%v, proxy=%v, interface=%v", configureOnly, manageProxy, iface)
	if !configureOnly {
		// No changes if we're not configuring the app
		return
	}

	// Update the settings
	cfg.NetworkInterface = iface
	cfg.ManageProxy = manageProxy
	err := config.StoreParameters(cfg)
	if err != nil {
		fmt.Println("Error saving parameters:", err)
		os.Exit(1)
	}
	fmt.Println("Device Config configured successfully")
	os.Exit(0)
}
