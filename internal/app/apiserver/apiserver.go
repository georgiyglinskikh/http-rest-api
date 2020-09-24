package apiserver

import (
	"io"       // Read | write strings through web and to files
	"net/http" // Work with http

	"github.com/gorilla/mux"     // Route library
	"github.com/sirupsen/logrus" // Logs for server
)

// APIServer type with all things that are necessary
type APIServer struct {
	Config *Config
	Logger *logrus.Logger
	Router *mux.Router
}

// New function for creating APIServer instance
func New(config *Config) *APIServer {
	return &APIServer{
		Config: config,
		Logger: logrus.New(),
		Router: mux.NewRouter(),
	}
}

// Start server & logging
func (server *APIServer) Start() error {
	// Init logger
	if err := server.configureLogger(); err != nil {
		return err
	}

	// Init router
	server.configureRouter()

	// Display info about server state
	server.Logger.Info("starting api server")

	// Run rest api with current instance of APIServer object
	return http.ListenAndServe(server.Config.BindAddress, server.Router)
}

// Init logger
func (server *APIServer) configureLogger() error {
	// Get logger level (debug | error) from file
	level, err := logrus.ParseLevel(server.Config.LogLevel)
	if err != nil {
		return err
	}

	// and set it
	server.Logger.SetLevel(level)

	return nil // Exit successfully
}

// Init router
func (server *APIServer) configureRouter() {
	// Map route with:
	server.Router.HandleFunc("/hello", server.handleHello())
}

// Router function that work with `/hello` address
func (server *APIServer) handleHello() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(response, "Hello") // Send string back
	}
}
