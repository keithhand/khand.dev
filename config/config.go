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

type config struct {
	Port int
}

func NewConfig() config {
	portEnv := os.Getenv(port.Key)
	if portEnv == "" {
		portEnv = port.Default
	}

	port, err := strconv.Atoi(portEnv)
	if err != nil {
		panic(err)
	}

	return config{
		Port: port,
	}
}
