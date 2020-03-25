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
	"github.com/CanonicalLtd/device-config/service/dbus"
	"testing"
)

func TestTime_Apply(t1 *testing.T) {
	type args struct {
		ntp      bool
		timezone string
		setTime  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"valid", args{false, "Europe/London", "2020-02-22T20:20:20Z"}, false},
		{"valid-time-sync", args{true, "Europe/London", ""}, false},
		{"invalid-timezone", args{true, "invalid", ""}, true},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewTime(&dbus.MockDbus{})
			if err := t.Apply(tt.args.ntp, tt.args.timezone, tt.args.setTime); (err != nil) != tt.wantErr {
				t1.Errorf("Apply() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTime_Current(t1 *testing.T) {
	t := NewTime(&dbus.MockDbus{})
	got := t.Current()

	if got == nil {
		t1.Errorf("Current() nil response")
		return
	}
	if len(got.Timezones) != len(dbus.Timezones) {
		t1.Errorf("Current() Timezones got = %v, want %v", len(got.Timezones), len(dbus.Timezones))
	}
	if got.Time.Hour() != 22 && got.Time.Minute() != 22 && got.Time.Second() != 22 {
		t1.Errorf("Current() time got = %v, want %v", got.Time, "22:22:22")
	}
	if !got.NTP {
		t1.Errorf("Current() NOT got = %v, want %v", got.NTP, true)
	}

}
