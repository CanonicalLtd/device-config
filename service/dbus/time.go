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
	"sort"
	"strings"
	"time"
)

// Time holds the dbus time settings
type Time struct {
	Timezone string
	NTP      bool
	Time     time.Time
}

// TimeDateStatus gets the current time settings on the device
func (db *DBus) TimeDateStatus() *Time {
	t := Time{}
	timedate1 := db.getBusObject("org.freedesktop.timedate1", "/org/freedesktop/timedate1")

	timeZone, err := timedate1.GetProperty("org.freedesktop.timedate1.Timezone")
	if err == nil {
		t.Timezone = strings.Trim(timeZone.String(), "\"")
	}
	ntp, err := timedate1.GetProperty("org.freedesktop.timedate1.NTP")
	if err == nil {
		t.NTP = ntp.Value().(bool)
	}
	timeUsec, err := timedate1.GetProperty("org.freedesktop.timedate1.TimeUSec")
	if err == nil {
		uu := timeUsec.Value().(uint64)
		t.Time = time.Unix(int64(uu/1e6), 0)
	}
	return &t
}

// SetNTP sets whether the time should be synced
func (db *DBus) SetNTP(value bool) error {
	// Set to use the NTP
	timedate1 := db.getBusObject("org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	call := timedate1.Call("org.freedesktop.timedate1.SetNTP", 0, value, false)
	return call.Err
}

// SetTimezone sets the device time zone
func (db *DBus) SetTimezone(timezone string) error {
	// Check we have a valid time zone
	i := sort.Search(len(Timezones), func(i int) bool { return Timezones[i] >= timezone })
	if i >= len(Timezones) || Timezones[i] != timezone {
		return fmt.Errorf("`%s` is not a valid time zone", timezone)
	}

	// Set the time zone
	timedate1 := db.getBusObject("org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	call := timedate1.Call("org.freedesktop.timedate1.SetTimezone", 0, timezone, false)
	return call.Err
}

// SetTime sets the current time
func (db *DBus) SetTime(setTime string) error {
	parsed, err := time.Parse("2006-01-02T15:04:05Z", setTime)
	if err != nil {
		return err
	}

	// Turn off time sync first
	if err := db.SetNTP(false); err != nil {
		return err
	}

	// Set the time
	timedate1 := db.getBusObject("org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	call := timedate1.Call("org.freedesktop.timedate1.SetTime", 0, parsed.UnixNano()/1000, false, false)
	return call.Err
}
