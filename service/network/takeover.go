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

package network

import (
	"gopkg.in/yaml.v2"
	"os"
)

const snapdFile = "/etc/netplan/00-snapd-config.yaml"

// TakeOver takes over network configuration for the device by removing
// existing netplan files
func TakeOver() error {
	// Check if we have already configured this device
	if _, err := os.Stat(netplanFilePath); err == nil {
		// Remove the snapd/console-conf netplan config file
		_ = os.Remove(snapdFile)
		return nil
	}

	// Check if the snapd netplan file is there
	if _, err := os.Stat(snapdFile); err != nil {
		// No - all done
		return nil
	}

	// Rename the snapd/console-conf file, if it's there
	// So we keep the existing config
	return os.Rename(snapdFile, netplanFilePath)
}

func serializeNetplan(deviceNetplan *NetplanYAML) error {
	// Serialize the data to YAML
	data, err := yaml.Marshal(deviceNetplan)
	if err != nil {
		return nil
	}

	// Write the YAML to the config file
	return writeNetplan(data)
}
