package apiserver

import (
	"net/http"          // Http work
	"net/http/httptest" // Test for network (http)
	"testing"           // Test utilities

	"github.com/stretchr/testify/assert" // Assert for testing
)

func TestAPIServer_HandleHello(t *testing.T) {
	server := NewAPIServer(DefaultConfig())                      // Create server
	response := httptest.NewRecorder()                           // Listen to server output
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil) // Send request
	server.handleHello().ServeHTTP(response, request)            // Send request for reading response
	assert.Equal(t, response.Body.String(), "Hello")             // Test if response matches expected string
}
