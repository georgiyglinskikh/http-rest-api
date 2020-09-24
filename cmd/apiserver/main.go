package main

import (
	"flag" // Parse cli flags
	"log"  // Log server initileasing and stopping

	"github.com/BurntSushi/toml"                                                // Parse toml
	apiserver "github.com/georgiyglinskikh/http-rest-api/iternal/app/apiserver" // Iternal logic
)

var (
	configPath string // Path to config files
)

// init Work before server starts
func init() {
	// Extract flag option to global variable
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// main Entery point for program
func main() {
	flag.Parse()

	// Create config for next using
	config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	// Create and start server
	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
