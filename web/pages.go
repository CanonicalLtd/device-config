// Ubuntu Core Configuration
// Copyright 2020 Canonical Ltd.  All rights reserved.

package web

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type commonData struct {
	Username string
	Error    string
}

func (srv Web) templates(name string) (*template.Template, error) {
	// Parse the templates
	p := filepath.Join(srv.Settings.DocRoot, name)
	t, err := template.ParseFiles(p)
	if err != nil {
		log.Printf("Error loading the application template: %v\n", err)
	}
	return t, err
}

func getUsername(r *http.Request) string {
	username, err := r.Cookie("username")
	if err != nil {
		return ""
	}
	return username.String()
}
