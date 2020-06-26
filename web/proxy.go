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
	"encoding/json"
	"io"
	"net/http"
)

// ProxyConfig holds the config for proxy servers
type ProxyConfig struct {
	HTTP  string `json:"http,omitempty"`
	HTTPS string `json:"https,omitempty"`
	FTP   string `json:"ftp,omitempty"`
}

// Proxy is the API to get the proxy config
func (srv Web) Proxy(w http.ResponseWriter, r *http.Request) {
	// Get the current settings
	systemConf, err := srv.Snapd.Conf("system")
	if err != nil {
		formatStandardResponse("proxy", err.Error(), w)
		return
	}

	// Return the response - snapd returns in ProxyConfig format
	formatProxyResponse(systemConf["proxy"], w)
}

// ProxyUpdate is the API to update the proxy config
func (srv Web) ProxyUpdate(w http.ResponseWriter, r *http.Request) {
	req := srv.decodeProxyConfig(w, r)
	if req == nil {
		return
	}

	// Set up the proxy settings
	if err := srv.Snapd.SetProxy(req.HTTP, req.HTTPS, req.FTP); err != nil {
		formatStandardResponse("proxy-update", err.Error(), w)
		return
	}
	formatStandardResponse("", "", w)
}

func (srv Web) decodeProxyConfig(w http.ResponseWriter, r *http.Request) *ProxyConfig {
	// Decode the JSON body
	req := ProxyConfig{}
	err := json.NewDecoder(r.Body).Decode(&req)
	switch {
	// Check we have some data
	case err == io.EOF:
		formatStandardResponse("data", "No proxy data supplied.", w)
		return nil
		// Check for parsing errors
	case err != nil:
		formatStandardResponse("decode-json", err.Error(), w)
		return nil
	}
	return &req
}
