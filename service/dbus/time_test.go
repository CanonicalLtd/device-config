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
	"testing"
)

func getMockBusObject(systemBus interface{}, dest, path string) dbus.BusObject {
	return &mockBusObject{}
}

func TestDBus_TimeDateStatus(t *testing.T) {
	busObject = getMockBusObject
	db := &DBus{systemBus: nil}

	got := db.TimeDateStatus()
	if got.Timezone != "Europe/London" {
		t.Errorf("TimeDateStatus() timezone = %v, want %v", got.Timezone, "Europe/London")
	}
}

func TestDBus_SetNTPTimezoneTime(t *testing.T) {
	busObject = getMockBusObject
	db := &DBus{systemBus: nil}

	if err := db.SetNTP(true); err != nil {
		t.Errorf("SetNTP() error = %v", err)
	}

	if err := db.SetTimezone("America/Barbados"); err != nil {
		t.Errorf("SetTimezone() error = %v", err)
	}
	if err := db.SetTimezone("America/OverTheRainbow"); err.Error() != "`America/OverTheRainbow` is not a valid time zone" {
		t.Errorf("SetTimezone() expected error = %v", "`America/OverTheRainbow` is not a valid time zone")
	}

	if err := db.SetTime("2020-01-02T15:04:05Z"); err != nil {
		t.Errorf("SetTime() error = %v", err)
	}
	if err := db.SetTime("outside time"); err == nil {
		t.Errorf("SetTime() expected error for `%v`", "outside time")
	}
}
