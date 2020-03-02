/*
 * Ubuntu Core Configuration
 * Copyright 2020 Canonical Ltd.  All rights reserved.
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

func encodeResponse(w http.ResponseWriter, response interface{}) {
	// Encode the response as JSON
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error forming the response:", err)
	}
}
