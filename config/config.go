// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package config

// Default settings
const (
	DefaultInterface     = "0.0.0.0"
	DefaultPort          = "8888"
	DefaultDocRoot       = "./static"
	DefaultIndexTemplate = "index.html"
)

// Settings defines the application configuration
type Settings struct {
	NetworkInterface string
	Port             string
	DocRoot          string
	IndexTemplate    string
}

// ParseArgs checks the environment variables
func ParseArgs() *Settings {
	return &Settings{
		NetworkInterface: DefaultInterface,
		Port:             DefaultPort,
		DocRoot:          DefaultDocRoot,
		IndexTemplate:    DefaultIndexTemplate,
	}
}
