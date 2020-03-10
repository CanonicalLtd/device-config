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

package service

import (
	"fmt"
	"time"
)

type mockDbus struct{}

func (m *mockDbus) TimeDateStatus() *DBusTime {
	parsed, _ := time.Parse("2006-01-02T15:04:05Z", "2020-02-22T22:22:22Z")
	return &DBusTime{
		Timezone: "America/Barbados",
		NTP:      true,
		Time:     parsed,
	}
}
func (m *mockDbus) SetNTP(value bool) error {
	return nil
}
func (m *mockDbus) SetTime(setTime string) error {
	_, err := time.Parse("2006-01-02T15:04:05Z", setTime)
	return err
}
func (m *mockDbus) NetplanApply() error {
	return nil
}

func (m *mockDbus) SetTimezone(timezone string) error {
	for t := range timezones {
		if timezones[t] == timezone {
			return nil
		}
	}
	return fmt.Errorf("MOCK error in timezone")
}

var readNetplanFileError = func() ([]byte, error) {
	return nil, fmt.Errorf("MOCK error reading netplan file")
}
var readNetplanFileSuccess = func() ([]byte, error) {
	return []byte(`
network:
  version: 2
  renderer: networkd
  ethernets:
    enp3s0:
      dhcp4: true`), nil
}
var readNetplanFileBad = func() ([]byte, error) {
	return []byte("\u1000"), nil
}
