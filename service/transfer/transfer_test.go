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

package transfer

import (
	"fmt"
	"github.com/CanonicalLtd/device-config/service/dbus"
	"github.com/CanonicalLtd/device-config/service/snapd"
	"github.com/snapcore/snapd/client"
	"reflect"
	"testing"
	"time"
)

type mockSnapd struct{}

func (snap *mockSnapd) AppServices(names []string) ([]*client.AppInfo, error) {
	panic("implement me")
}

func (snap *mockSnapd) Conf(name string) (map[string]interface{}, error) {
	if name == "system" {
		proxy := map[string]string{"https": "https://proxy", "http": "http://proxy", "ftp": "ftp://proxy"}
		return map[string]interface{}{"proxy": proxy}, nil
	}
	return nil, fmt.Errorf("MOCK conf error")
}

func (snap *mockSnapd) List(names []string, opts *client.ListOptions) ([]snapd.Snap, error) {
	panic("implement me")
}

func (snap *mockSnapd) SetConf(name string, patch map[string]interface{}) (string, error) {
	panic("implement me")
}

func (snap *mockSnapd) SetProxy(http, https, ftp string) error {
	return nil
}

type mockDbus struct{}

func (db *mockDbus) TimeDateStatus() *dbus.Time {
	return &dbus.Time{
		Timezone: "Europe/London",
		NTP:      true,
		Time:     time.Now(),
	}
}

func (db *mockDbus) SetNTP(value bool) error {
	return nil
}

func (db *mockDbus) SetTimezone(timezone string) error {
	return nil
}

func (db *mockDbus) SetTime(setTime string) error {
	panic("implement me")
}

func (db *mockDbus) NetplanApply() error {
	panic("implement me")
}

func (db *mockDbus) NMIsRunning() error {
	panic("implement me")
}

func (db *mockDbus) NMDevices() (map[string]string, error) {
	panic("implement me")
}

func (db *mockDbus) NMInterfaceConfig(p string) *dbus.NMDeviceSettings {
	panic("implement me")
}

func (db *mockDbus) NMInterfaceConfigUpdate(p string, eth dbus.NMDeviceSettings) error {
	panic("implement me")
}

func TestTransfer_Export(t *testing.T) {
	valid := &Config{
		NTP:        true,
		Timezone:   "Europe/London",
		ProxyHTTP:  "http://proxy",
		ProxyHTTPS: "https://proxy",
		ProxyFTP:   "ftp://proxy",
	}

	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{"valid", valid, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewTransfer(&mockDbus{}, &mockSnapd{})

			got, err := x.Export()
			if (err != nil) != tt.wantErr {
				t.Errorf("Export() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Export() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransfer_Import(t *testing.T) {
	valid := Config{
		NTP:        true,
		Timezone:   "Europe/London",
		ProxyHTTP:  "http://proxy",
		ProxyHTTPS: "https://proxy",
		ProxyFTP:   "ftp://proxy",
	}

	type args struct {
		cfg Config
	}
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{"valid", valid, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := NewTransfer(&mockDbus{}, &mockSnapd{})
			if err := x.Import(tt.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Import() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
