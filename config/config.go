package config

import (
	"log"
	"os"
	"strconv"
)

var (
	ServerPort = NewEnv("SERVER_PORT", 8080)
	GHProfile  = NewEnv("GH_PROFILE", "keithhand")
)

type EnvTypes interface {
	int | string
}

func NewEnv[T EnvTypes](key string, defaultValue T) T {
	val := defaultValue

	defer func() {
		log.Printf("finished configuring envvar: %s:%v\n", key, val)
	}()

	env := os.Getenv(key)
	if env == "" {
		return val
	}

	switch vp := any(&val).(type) {
	case *string:
		*vp = env
	case *int:
		ival, err := strconv.Atoi(env)
		if err != nil {
			log.Printf("error reading env value %s. got %s expected int. falling back to %v", key, env, val)
			return val
		}
		*vp = ival
	}

	return val
}
