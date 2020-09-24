package apiserver

import (
	"net/http"          // Http work
	"net/http/httptest" // Test for network (http)
	"testing"           // Test utilities

	"github.com/stretchr/testify/assert" // Assert for testing
)

func TestAPIServer_HandleHello(t *testing.T) {
	server := New(NewConfig())                                   // Create server
	responce := httptest.NewRecorder()                           // Listen to server output
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil) // Send request
	server.handleHello().ServeHTTP(responce, request)            // Send request for reading responce
	assert.Equal(t, responce.Body.String(), "Hello")             // Test if responce matches expected string
}
