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

package web

import (
	"encoding/base64"
	"github.com/CanonicalLtd/device-config/service/transfer"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// TransferExport generates a file to export the current configuration
func (srv Web) TransferExport(w http.ResponseWriter, r *http.Request) {
	cfg, err := srv.TransferSrv.Export()
	if err != nil {
		formatStandardResponse("transfer", err.Error(), w)
		return
	}

	// Convert to YAML
	data, err := yaml.Marshal(cfg)
	if err != nil {
		formatStandardResponse("transfer", err.Error(), w)
		return
	}

	// Base64 encode the data and submit it
	dataEnc := base64.StdEncoding.EncodeToString(data)
	io.Copy(w, strings.NewReader(dataEnc))
}

// TransferImport gets an import configuration file and configures the current system
func (srv Web) TransferImport(w http.ResponseWriter, r *http.Request) {
	// Decode the message body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		formatStandardResponse("transfer", err.Error(), w)
		return
	}

	// Decode the base64-encoded data
	data, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		formatStandardResponse("transfer", err.Error(), w)
		return
	}

	// Decode the YAML data
	cfg := transfer.Config{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		formatStandardResponse("transfer", err.Error(), w)
		return
	}

	// Import the settings
	if err := srv.TransferSrv.Import(cfg); err != nil {
		formatStandardResponse("transfer", err.Error(), w)
		return
	}

	formatStandardResponse("", "", w)
}
