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
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Custom holds the custom settings from the content interface
type Custom struct {
	Title    string `yaml:"title" json:"title"`
	SubTitle string `yaml:"subtitle" json:"subtitle"`

	Copyright string              `yaml:"copyright" json:"copyright"`
	Bullets   []map[string]string `yaml:"bullet" json:"bullet"`
}

func readCustomSettings() Custom {
	config := Custom{}

	p := path.Join(os.Getenv("SNAP"), "static", "custom", "custom.yaml")
	dat, err := ioutil.ReadFile(p)
	if err != nil {
		log.Printf("Error reading config parameters: %v", err)
		return defaultCustom()
	}

	if err := yaml.Unmarshal(dat, config); err != nil {
		return defaultCustom()
	}
	return config
}

func defaultCustom() Custom {
	return Custom{
		Title:     "Ubuntu Core Configuration",
		SubTitle:  "Configuration service for Ubuntu Core devices.",
		Copyright: "© 2020 Canonical Ltd. Ubuntu and Canonical are registered trademarks of Canonical Ltd.",
		Bullets: []map[string]string{
			{"text": "Legal information", "url": "https://ubuntu.com/legal"},
			{"text": "Privacy", "url": "https://ubuntu.com/legal/data-privacy"},
			{"text": "Report a bug on this site", "url": "https://github.com/CanonicalLtd/device-config/issues/new"},
		},
	}
}
