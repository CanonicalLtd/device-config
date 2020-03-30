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
	"fmt"
	"time"
)

// MockDbus mocks dbus
type MockDbus struct{}

// TimeDateStatus mocks the time/date config
func (m *MockDbus) TimeDateStatus() *Time {
	parsed, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-22T22:22:22Z")
	return &Time{
		Timezone: "America/Barbados",
		NTP:      true,
		Time:     parsed,
	}
}

// SetNTP mocks the ntp server usages
func (m *MockDbus) SetNTP(value bool) error {
	return nil
}

// SetTime mocks setting time
func (m *MockDbus) SetTime(setTime string) error {
	_, err := time.Parse("2006-01-02T15:04:05Z", setTime)
	return err
}

// NetplanApply mocks applying netplan config
func (m *MockDbus) NetplanApply() error {
	return nil
}

// NMIsRunning mocks checking network manager
func (m *MockDbus) NMIsRunning() error {
	return nil
}

//NMInterfaceConfig mocks the network manager config
func (m *MockDbus) NMInterfaceConfig(p string) *NMDeviceSettings {
	return &NMDeviceSettings{
		State:       100,
		DHCP4:       false,
		AddressData: []NMDeviceAddress{{"192.168.1.100", 24}},
		NameServers: []string{"192.168.1.1", "8.8.8.8"},
		Gateway:     "",
	}
}

// NMDevices mocks fetching the configured devices
func (m *MockDbus) NMDevices() (map[string]string, error) {
	return map[string]string{}, nil
}

//SetTimezone mocks setting the time zone
func (m *MockDbus) SetTimezone(timezone string) error {
	for t := range Timezones {
		if Timezones[t] == timezone {
			return nil
		}
	}
	return fmt.Errorf("MOCK error in timezone")
}

// NMInterfaceConfigUpdate updates network settings
func (m *MockDbus) NMInterfaceConfigUpdate(p string, eth NMDeviceSettings) error {
	return nil
}
