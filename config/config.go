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
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Default settings
const (
	DefaultInterface     = "0.0.0.0"
	DefaultPort          = "8888"
	DefaultDocRoot       = "./static"
	DefaultIndexTemplate = "index.html"
	DefaultManageProxy   = false
	paramsEnvVar         = "SNAP_DATA"
	paramsFilename       = "params"
)

// Settings defines the application configuration
type Settings struct {
	NetworkInterface string
	Port             string
	DocRoot          string
	IndexTemplate    string
	ManageProxy      bool
}

// DefaultArgs checks the environment variables
func DefaultArgs() *Settings {
	return &Settings{
		NetworkInterface: DefaultInterface,
		Port:             DefaultPort,
		DocRoot:          DefaultDocRoot,
		IndexTemplate:    DefaultIndexTemplate,
		ManageProxy:      DefaultManageProxy,
	}
}

// ReadParameters fetches the store config parameters
func ReadParameters() *Settings {
	config := DefaultArgs()

	p := GetPath(paramsFilename)

	dat, err := ioutil.ReadFile(p)
	if err != nil {
		log.Printf("Error reading config parameters: %v", err)
		return config
	}

	err = json.Unmarshal(dat, config)
	if err != nil {
		log.Printf("Error parsing config parameters: %v", err)
		return config
	}

	return config
}

// StoreParameters stores the configuration parameters on the filesystem
func StoreParameters(c *Settings) error {
	p := GetPath(paramsFilename)

	// Default empty parameters
	if len(c.NetworkInterface) == 0 {
		c.NetworkInterface = DefaultInterface
	}
	if len(c.Port) == 0 {
		c.Port = DefaultPort
	}
	if len(c.DocRoot) == 0 {
		c.DocRoot = DefaultDocRoot
	}
	if len(c.IndexTemplate) == 0 {
		c.IndexTemplate = DefaultIndexTemplate
	}

	log.Println("---Store:", c)

	// Create the output file
	f, err := os.Create(p)
	if err != nil {
		return err
	}
	defer f.Close()

	// Convert the parameters to JSON
	b, err := json.Marshal(c)
	if err != nil {
		log.Printf("Error marshalling config parameters: %v", err)
		return err
	}

	// Output the JSON to the file
	_, err = f.Write(b)
	if err != nil {
		log.Printf("Error storing config parameters: %v", err)
		return err
	}
	_ = f.Sync()

	// Restrict access to the file
	return os.Chmod(p, 0600)
}

// GetPath returns the full path to the data file
func GetPath(filename string) string {
	if len(os.Getenv(paramsEnvVar)) > 0 {
		return path.Join(os.Getenv(paramsEnvVar), "../current", filename)
	}
	return filename
}
