# http-rest-api

Simple rest api application for learning purposes

It's written entirely in `GO`

## Preservatives

- github.com/BurntSushi/toml = Parse toml with ease
- github.com/gorilla/mux = Library for routing
- github.com/sirupsen/logrus = Log server
- github.com/stretchr/testify = Library for tests

## Structure:
- `cmd/` - public implementation
- `configs/` - config files (in toml)
- `iternal/` - all logic and realisations
- `go.mod` - dependencies
- `Makefile` - file with all modes to build, run and test
- `README.md` - file you're reading right now

## Build from source

### You need:
- make
- go (1.52+)

### Build
`$ make`

### Run tests
`$ make test`