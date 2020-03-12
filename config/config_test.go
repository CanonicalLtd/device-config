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
	tests := []struct {
		name string
		want *Settings
	}{
		{"valid", &Settings{DefaultInterface, DefaultPort, DefaultDocRoot, DefaultIndexTemplate, DefaultManageProxy}},
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
	if got.ManageProxy != DefaultManageProxy {
		t.Errorf("ReadParameters() got = %v, want %v", got.ManageProxy, DefaultManageProxy)
	}
	if got.NetworkInterface != DefaultInterface {
		t.Errorf("ReadParameters() got = %v, want %v", got.NetworkInterface, DefaultInterface)
	}

	_ = os.Remove(paramsFilename)
}

func TestStoreParameters(t *testing.T) {
	cfg := &Settings{}
	err := StoreParameters(cfg)
	if err != nil {
		t.Errorf("StoreParameters() error = %v", err)
	}
	if cfg.ManageProxy != DefaultManageProxy {
		t.Errorf("ReadParameters() got = %v, want %v", cfg.ManageProxy, DefaultManageProxy)
	}
	if cfg.NetworkInterface != DefaultInterface {
		t.Errorf("ReadParameters() got = %v, want %v", cfg.NetworkInterface, DefaultInterface)
	}

	_ = os.Remove(paramsFilename)
}
