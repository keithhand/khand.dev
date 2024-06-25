package config

import (
	"os"
	"strconv"
)

var (
	port = envVar{
		Key:     "APP_PORT",
		Default: "8080",
	}
)

type envVar struct {
	Key     string
	Default string
}

type appConfig struct {
	Port int
}

func New() appConfig {
	portEnv := os.Getenv(port.Key)
	if portEnv == "" {
		portEnv = port.Default
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}

	return appConfig{
		Port: port,
	}
}
