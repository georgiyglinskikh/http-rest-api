.PHONEY: build
build: # Build app
	go build -v ./cmd/apiserver

.PHONEY: test
test: # Run tests
	go test -v -race -timeout 30s ./...
	
.DEFAULT_GOAL := build