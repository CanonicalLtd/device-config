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

package main

import (
	"flag"
	"fmt"
	"github.com/CanonicalLtd/device-config/config"
	"os"
)

func main() {
	var (
		iface       string
		manageProxy bool
	)
	flag.StringVar(&iface, "interface", config.DefaultInterface, "The default network interface for the service")
	flag.BoolVar(&manageProxy, "proxy", config.DefaultManageProxy, "Allow proxy configuration (needs the snapd-control interface)")
	flag.Parse()

	// Read the config settings
	cfg := config.ReadParameters()

	// Update the settings
	cfg.NetworkInterface = iface
	cfg.ManageProxy = manageProxy
	err := config.StoreParameters(cfg)
	if err != nil {
		fmt.Println("Error saving parameters:", err)
		os.Exit(1)
	}
}
