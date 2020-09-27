package apiserver

import "flag" // Parse cli flags

var (
	ConfigPath string
)

func ExtractFlagsToGlobalVariables() {
	extractConfigPathFlag()

	flag.Parse()
}

func extractConfigPathFlag() {
	flag.StringVar(
		&ConfigPath,
		"config-path",
		"configs/apiserver.toml",
		"path to config file",
	)
}
