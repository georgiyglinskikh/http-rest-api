package main

import (
	. "github.com/georgiyglinskikh/http-rest-api/internal/app/apiserver"
)

func main() {
	ExtractFlagsToGlobalVariables()

	config := FromTOML(ConfigPath)

	server := NewAPIServer(config)
	server.Run()
}
