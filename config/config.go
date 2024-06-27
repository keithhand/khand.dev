package config

import (
	"fmt"
	"os"
	"strconv"

	"khand.dev/khand.dev/logs"
)

type Config struct {
	GHProfile  string
	ServerPort int
}

func New() *Config {
	return &Config{
		GHProfile:  newEnv("GH_PROFILE", "keithhand"),
		ServerPort: newEnv("SERVER_PORT", 8080),
	}
}

type EnvTypes interface {
	int | string
}

func newEnv[T EnvTypes](key string, defaultValue T) T {
	val := defaultValue
	defer func() {
		logs.Debug("finished configuring env:", key, val)
	}()

	env := os.Getenv(key)
	if env == "" {
		logs.Debug("env value not set, using default", key, val)
		return val
	}

	if err := setValueFromEnv(&val, env); err != nil {
		logs.Error(fmt.Errorf("config: setting %s: %w", key, err).Error())
	}

	return val
}

func setValueFromEnv[T EnvTypes](val *T, env string) error {
	switch vp := any(val).(type) {

	case *string:
		*vp = env

	case *int:
		ival, err := strconv.Atoi(env)
		if err != nil {
			return fmt.Errorf("converting %s to %T", env, *val)
		}
		*vp = ival

	default:
		return fmt.Errorf("type not found: %T expected, got %s", *val, env)
	}
	return nil
}
