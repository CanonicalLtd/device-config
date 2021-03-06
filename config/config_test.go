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

package config

import (
	"os"
	"reflect"
	"testing"
)

func TestDefaultArgs(t *testing.T) {
	cfg := &Settings{
		NetworkInterfaceDevice: DefaultInterfaceDevice,
		NetworkInterfaceIP:     DefaultInterfaceIP,
		Port:                   DefaultPort,
		DocRoot:                DefaultDocRoot,
		IndexTemplate:          DefaultIndexTemplate,
		SnapControl:            DefaultSnapControl,
		UseNetworkManager:      DefaultUseNetworkManager,
		HideInterfaces:         []string{},
	}

	tests := []struct {
		name string
		want *Settings
	}{
		{"valid", cfg},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultArgs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadParameters(t *testing.T) {
	got := ReadParameters()
	if got.SnapControl != DefaultSnapControl {
		t.Errorf("ReadParameters() got = %v, want %v", got.SnapControl, DefaultSnapControl)
	}
	if got.NetworkInterfaceIP != DefaultInterfaceIP {
		t.Errorf("ReadParameters() got = %v, want %v", got.NetworkInterfaceIP, DefaultInterfaceIP)
	}

	_ = os.Remove(paramsFilename)
}

func TestStoreParameters(t *testing.T) {
	cfg := &Settings{}
	err := StoreParameters(cfg)
	if err != nil {
		t.Errorf("StoreParameters() error = %v", err)
	}
	if cfg.SnapControl != DefaultSnapControl {
		t.Errorf("ReadParameters() got = %v, want %v", cfg.SnapControl, DefaultSnapControl)
	}
	if cfg.NetworkInterfaceIP != DefaultInterfaceIP {
		t.Errorf("ReadParameters() got = %v, want %v", cfg.NetworkInterfaceIP, DefaultInterfaceIP)
	}

	_ = os.Remove(paramsFilename)
}
