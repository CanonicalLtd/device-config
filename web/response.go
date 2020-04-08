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
	"log"
	"net/http"
)

// JSONHeader is the header for JSON responses
const JSONHeader = "application/json; charset=UTF-8"

// StandardResponse is the JSON response from an API method, indicating success or failure.
type StandardResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RecordsResponse the JSON response from a list call
type RecordsResponse struct {
	StandardResponse
	Records interface{} `json:"records"`
}

// LoginResponse is the response message from a login action
type LoginResponse struct {
	StandardResponse
	Username  string `json:"username"`
	SessionID string `json:"sessionId"`
}

// NetworkResponse the JSON response from a network call
type NetworkResponse struct {
	StandardResponse
	Interfaces []InterfaceConfig `json:"interfaces"`
}

// ProxyResponse the JSON response from a proxy config call
type ProxyResponse struct {
	StandardResponse
	Proxy interface{} `json:"proxy"`
}

// TimeResponse the JSON response from a time config call
type TimeResponse struct {
	StandardResponse
	Time interface{} `json:"time"`
}

// ServiceResponse the JSON response from a app services call
type ServiceResponse struct {
	StandardResponse
	Services interface{} `json:"services"`
}

// AppConfigResponse the JSON response from a app config call
type AppConfigResponse struct {
	StandardResponse
	Services interface{} `json:"config"`
}

// formatStandardResponse returns a JSON response from an API method, indicating success or failure
func formatStandardResponse(code, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := StandardResponse{Code: code, Message: message}

	if len(code) > 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatLoginResponse returns a JSON response from a login
func formatLoginResponse(username, sessionID string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := LoginResponse{StandardResponse{}, username, sessionID}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatNetworkResponse returns a JSON response from a network call
func formatNetworkResponse(interfaces []InterfaceConfig, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := NetworkResponse{StandardResponse{}, interfaces}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatProxyResponse returns a JSON response from a proxy call
func formatProxyResponse(proxy interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := ProxyResponse{StandardResponse{}, proxy}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatTimeResponse returns a JSON response from a time config call
func formatTimeResponse(cfg interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := TimeResponse{StandardResponse{}, cfg}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatServicesResponse returns a JSON response from an app services call
func formatServicesResponse(status interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := ServiceResponse{StandardResponse{}, status}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatAppConfigResponse returns a JSON response from an app services call
func formatAppConfigResponse(cfg interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := AppConfigResponse{StandardResponse{}, cfg}

	// Encode the response as JSON
	encodeResponse(w, response)
}

// formatRecordsResponse returns a JSON response from an app services call
func formatRecordsResponse(records interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", JSONHeader)
	response := RecordsResponse{StandardResponse{}, records}

	// Encode the response as JSON
	encodeResponse(w, response)
}

func encodeResponse(w http.ResponseWriter, response interface{}) {
	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error forming the response:", err)
	}
}
