/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
 *
 */

package service

import (
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
	Time      time.Time `json:"time"`

	dBus DBusService `json:"-"`
}

// NewTime creates a time object from the device settings
func NewTime(dBus DBusService) *Time {
	return &Time{
		Timezones: timezones,
		dBus:      dBus,
	}
}

// Current returns the current time settings
func (t *Time) Current() *Time {
	dbusTime := t.dBus.TimeDateStatus()
	t.Time = dbusTime.Time
	t.Timezone = dbusTime.Timezone
	t.NTP = dbusTime.NTP

	return t
}

// Apply updates the time settings
func (t *Time) Apply(ntp bool, timezone, setTime string) error {
	if err := t.dBus.SetTimezone(timezone); err != nil {
		return nil
	}

	// Set up time sync
	if ntp {
		return t.dBus.SetNTP(true)
	}

	// Manually set the time
	return t.dBus.SetTime(setTime)
}
