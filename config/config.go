package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server
	GitHubApi
}

var log logger

type logger interface {
	Debug(string, ...any)
	Error(string, ...any)
}

func New(lgr logger) *Config {
	log = lgr
	return &Config{
		GitHubApi: GitHubApi{
			profile: newEnv("GH_PROFILE", "keithhand"),
		},
		Server: Server{
			port: newEnv("SERVER_PORT", 8080),
		},
	}
}

type Server struct {
	port int
}

func (cfg Server) Port() int {
	return cfg.port
}

type GitHubApi struct {
	profile string
}

func (cfg GitHubApi) GHProfile() string {
	return cfg.profile
}

type EnvTypes interface {
	int | string
}

func newEnv[T EnvTypes](key string, def T) T {
	val := def
	defer func() {
		log.Debug("finished configuring env:", key, val)
	}()

	env := os.Getenv(key)
	if env == "" {
		log.Debug("env value not set, falling back to default", key, val)
		return val
	}

	if err := setValueFromEnv(&val, env); err != nil {
		log.Error(fmt.Errorf("config: setting %s: %w", key, err).Error())
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
