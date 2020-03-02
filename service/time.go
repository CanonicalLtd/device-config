/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

package service

import (
	"fmt"
	"github.com/godbus/dbus"
	"log"
	"sort"
	"strings"
	"time"
)

const timeDateCtl = "timedatectl"

// TimeService is the interface for the timedatectl command
type TimeService interface {
	Current() *Time
	Apply(ntp bool, timezone, setTime string) error
}

// Time implements actions for managing time
type Time struct {
	Timezones []string  `json:"timezones"`
	Timezone  string    `json:"timezone"`
	NTP       bool      `json:"ntp"`
	NTPSync   bool      `json:"ntpSync"`
	Time      time.Time `json:"time"`
	TimeRTC   time.Time `json:"timeRTC"`

	systemBus *dbus.Conn `json:"-"`
}

// NewTime creates a time object from the device settings
func NewTime() (*Time, error) {
	bus, err := dbus.SystemBus()
	if err != nil {
		log.Printf("Failed to access system dbus: %v", err)
		return nil, err
	}
	return &Time{
		Timezones: timezones,
		systemBus: bus,
	}, nil
}

// Current returns the current time settings
func (t *Time) Current() *Time {
	t.timeStatus()
	return t
}

// Apply updates the time settings
func (t *Time) Apply(ntp bool, timezone, setTime string) error {
	if err := t.setTimezone(timezone); err != nil {
		return nil
	}

	// Set up time sync
	if ntp {
		return t.setNTP(true)
	}

	// Manually set the time
	return t.setTime(setTime)
}

func (t *Time) setNTP(value bool) error {
	// Set to use the NTP
	timedate1 := t.systemBus.Object("org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	call := timedate1.Call("org.freedesktop.timedate1.SetNTP", 0, value, false)
	return call.Err
}

func (t *Time) setTimezone(timezone string) error {
	// Check we have a valid time zone
	i := sort.Search(len(timezones), func(i int) bool { return timezones[i] >= timezone })
	if i >= len(timezones) || timezones[i] != timezone {
		return fmt.Errorf("`%s` is not a valid time zone", timezone)
	}

	// Set the time zone
	timedate1 := t.systemBus.Object("org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	call := timedate1.Call("org.freedesktop.timedate1.SetTimezone", 0, timezone, false)
	return call.Err
}

func (t *Time) setTime(setTime string) error {
	parsed, err := time.Parse("2006-01-02T15:04:05Z", setTime)
	if err != nil {
		return err
	}

	// Turn off time sync first
	if err := t.setNTP(false); err != nil {
		return err
	}

	// Set the time
	timedate1 := t.systemBus.Object("org.freedesktop.timedate1", "/org/freedesktop/timedate1")
	call := timedate1.Call("org.freedesktop.timedate1.SetTime", 0, parsed.UnixNano()/1000, false, false)
	return call.Err
}

func (t *Time) timeStatus() {
	timedate1 := t.systemBus.Object("org.freedesktop.timedate1", "/org/freedesktop/timedate1")

	timeZone, err := timedate1.GetProperty("org.freedesktop.timedate1.Timezone")
	if err == nil {
		t.Timezone = strings.Trim(timeZone.String(), "\"")
	}
	ntp, err := timedate1.GetProperty("org.freedesktop.timedate1.NTP")
	if err == nil {
		t.NTP = ntp.Value().(bool)
	}
	ntpSync, err := timedate1.GetProperty("org.freedesktop.timedate1.NTPSynchronized")
	if err == nil {
		t.NTPSync = ntpSync.Value().(bool)
	}
	timeUsec, err := timedate1.GetProperty("org.freedesktop.timedate1.TimeUSec")
	if err == nil {
		uu := timeUsec.Value().(uint64)
		t.Time = time.Unix(int64(uu/1e6), 0)
	}
	rtcTimeUsec, err := timedate1.GetProperty("org.freedesktop.timedate1.RTCTimeUSec")
	if err == nil {
		uurtc := rtcTimeUsec.Value().(uint64)
		t.TimeRTC = time.Unix(int64(uurtc/1e6), 0)
	}
}
