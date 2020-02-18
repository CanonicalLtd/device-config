package main

import (
	"github.com/CanonicalLtd/configurator/config"
	"github.com/CanonicalLtd/configurator/web"
	"log"
)

func main() {
	// Parse the command-line arguments
	settings := config.ParseArgs()

	// Set up the dependency chain
	srv := web.NewWebService(settings)

	// Start the web service
	log.Fatal(srv.Start())
}
