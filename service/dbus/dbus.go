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

package dbus

import (
	"github.com/godbus/dbus"
	"log"
)

// Service is the interface to abstract dbus
type Service interface {
	TimeDateStatus() *Time
	SetNTP(value bool) error
	SetTimezone(timezone string) error
	SetTime(setTime string) error
	NetplanApply() error

	// Network Manager
	NMIsRunning() error
	NMDevices() (map[string]string, error)
	NMInterfaceConfig(p string) *NMDeviceSettings
	NMInterfaceConfigUpdate(p string, eth NMDeviceSettings) error
}

// DBus implements a wrapper for the dbus service
type DBus struct {
	systemBus *dbus.Conn
}

// NewDBus creates a dbus wrapper service
func NewDBus() (*DBus, error) {
	bus, err := dbus.SystemBus()
	if err != nil {
		log.Printf("Failed to access system dbus: %v", err)
		return nil, err
	}
	return &DBus{bus}, nil
}

// NetplanApply applies the current netplan configuration
func (db *DBus) NetplanApply() error {
	nPlan := db.getBusObject("io.netplan.NetSrv", "/io/netplan/NetSrv")
	call := nPlan.Call("io.netplan.NetSrv.Apply", 0)
	return call.Err
}

func (db *DBus) getBusObject(dest, path string) dbus.BusObject {
	return busObject(db.systemBus, dest, path)
}

// busObject returns a dbus bus object interface (mockable for tests)
var busObject = func(systemBus interface{}, dest, path string) dbus.BusObject {
	return systemBus.(*dbus.Conn).Object(dest, dbus.ObjectPath(path))
}
