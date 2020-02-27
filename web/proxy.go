/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
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
	settings := map[string]interface{}{}
	if len(req.HTTP) > 0 {
		settings["http"] = req.HTTP
	} else {
		settings["http"] = nil
	}
	if len(req.HTTPS) > 0 {
		settings["https"] = req.HTTPS
	} else {
		settings["https"] = nil
	}
	if len(req.FTP) > 0 {
		settings["ftp"] = req.FTP
	} else {
		settings["ftp"] = nil
	}

	cfg := map[string]interface{}{"proxy": settings}

	// Save the proxy settings
	_, err := srv.Snapd.SetConf("system", cfg)
	if err != nil {
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
