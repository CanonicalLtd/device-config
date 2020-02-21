// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"html/template"
	"log"
	"path/filepath"
)

func (srv Web) templates(name string) (*template.Template, error) {
	// Parse the templates
	p := filepath.Join(srv.Settings.DocRoot, name)
	pHeader := filepath.Join(srv.Settings.DocRoot, "header.html")
	pLayout := filepath.Join(srv.Settings.DocRoot, "layout.html")

	t, err := template.ParseFiles(pLayout, pHeader, p)
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
	}
	return t, err
}
