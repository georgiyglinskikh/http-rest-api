package apiserver

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	Config *Config
	Logger *logrus.Logger
	Router *mux.Router
}

func DefaultAPIServer() *APIServer {
	return &APIServer{
		Config: DefaultConfig(),
		Logger: logrus.New(),
		Router: mux.NewRouter(),
	}
}

func NewAPIServer(config *Config) *APIServer {
	server := DefaultAPIServer()
	server.Config = config
	return server
}

// Start server & logging
func (server *APIServer) Start() error {
	server.configure()

	server.Logger.Info("starting api server")

	// Run rest api with current instance of APIServer object
	return http.ListenAndServe(server.Config.BindAddress, server.Router)
}

func (server *APIServer) Run() {
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

func (server *APIServer) configure() {
	server.configureLogger()
	server.configureRouter()
}

func (server *APIServer) configureLogger() {
	level, err := logrus.ParseLevel(server.Config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	server.Logger.SetLevel(level)
}

func (server *APIServer) configureRouter() {
	server.Router.HandleFunc("/hello", server.handleHello())
}

func (server *APIServer) handleHello() http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(response, "Hello") // Send string back
	}
}
