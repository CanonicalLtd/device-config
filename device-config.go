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
	"github.com/CanonicalLtd/device-config/service/snapd"
	"github.com/CanonicalLtd/device-config/service/system"
	"github.com/CanonicalLtd/device-config/web"
	"log"
	"os"
	"strings"
)

func main() {
	// Get the application parameters
	settings := config.ReadParameters()
	configure(settings)

	// Set up the dependency chain
	memorySrv := memory.NewStore()
	authSrv := service.NewAuthService(memorySrv)
	snapdClient := snapd.NewClientAdapter()
	dBus, err := dbus.NewDBus()
	if err != nil {
		log.Fatal(err)
	}
	timeSrv := service.NewTime(dBus)
	systemSrv := system.NewSystem()
	netSrv := network.Factory(settings, dBus)

	srv := web.NewWebService(settings, authSrv, netSrv, snapdClient, timeSrv, systemSrv)

	// Start the web service
	log.Fatal(srv.Start())
}

func configure(cfg *config.Settings) {
	var (
		configureOnly bool
		iface         string
		listenOn      string
		snapControl   bool
		useNM         bool
		hideIfaces    string
	)
	flag.BoolVar(&configureOnly, "configure", false, "Configure the application and exit")
	flag.StringVar(&iface, "interface", config.DefaultInterfaceIP, "The default network interface for the service")
	flag.StringVar(&listenOn, "listenon", config.DefaultInterfaceDevice, "Force the service to listen a specific network device e.g. eth0")
	flag.BoolVar(&snapControl, "snapcontrol", config.DefaultSnapControl, "Display configuration that needs the snapd-control interface)")
	flag.BoolVar(&useNM, "nm", config.DefaultUseNetworkManager, "Use network manager instead of netplan")
	flag.StringVar(&hideIfaces, "hide", "", "Comma-separated list of interfaces to hide")
	flag.Parse()

	log.Printf("Device config: configure=%v, proxy=%v, interface=%v, nm=%v, hide=%v", configureOnly, snapControl, iface, useNM, hideIfaces)
	if !configureOnly {
		// No changes if we're not configuring the app
		return
	}

	// Update the settings
	cfg.NetworkInterfaceIP = iface
	cfg.NetworkInterfaceDevice = listenOn
	cfg.SnapControl = snapControl
	cfg.UseNetworkManager = useNM
	cfg.HideInterfaces = strings.Split(hideIfaces, ",")
	err := config.StoreParameters(cfg)
	if err != nil {
		fmt.Println("Error saving parameters:", err)
		os.Exit(1)
	}
	fmt.Println("Device Config configured successfully")
	os.Exit(0)
}
