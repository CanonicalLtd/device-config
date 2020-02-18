// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package config

// Default settings
const (
	DefaultPort          = "8888"
	DefaultDocRoot       = "./static"
	DefaultIndexTemplate = "index.html"
)

// Settings defines the application configuration
type Settings struct {
	Port          string
	DocRoot       string
	IndexTemplate string
}

// ParseArgs checks the environment variables
func ParseArgs() *Settings {
	return &Settings{
		Port:          ":" + DefaultPort,
		DocRoot:       DefaultDocRoot,
		IndexTemplate: DefaultIndexTemplate,
	}
}
