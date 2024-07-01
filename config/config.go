package config

import (
	"fmt"
	"os"
	"strconv"
)

type Logger interface {
	Debug(string, ...any)
	Error(string, ...any)
}

type Config struct {
	Server
	GitHubApi
}

type Server struct {
	port int
}

type GitHubApi struct {
	mockApi bool
	profile string
}

func New(lgr Logger) *Config {
	return &Config{
		GitHubApi: GitHubApi{
			mockApi: newEnv("MOCK_API", true, lgr),
			profile: newEnv("GH_PROFILE", "keithhand", lgr),
		},
		Server: Server{
			port: newEnv("SERVER_PORT", 8080, lgr),
		},
	}
}

func (cfg Server) Port() int {
	return cfg.port
}

func (cfg GitHubApi) MockApi() bool {
	return cfg.mockApi
}

func (cfg GitHubApi) GhProfile() string {
	return cfg.profile
}

type EnvTypes interface {
	int | string | bool
}

func newEnv[T EnvTypes](key string, def T, log Logger) T {
	fbMsg := "using default:"
	val := def
	defer func() {
		log.Debug("finished configuring env:", key, val)
	}()

	env := os.Getenv(key)
	if env == "" {
		log.Debug(fmt.Sprintf("env value not set, %s", fbMsg), key, val)
		return val
	}

	if err := setValueFromEnv(&val, env); err != nil {
		log.Error(
			fmt.Errorf("config: setting %s %w, %s", key, err, fbMsg).Error(),
			key, val)
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
			return fmt.Errorf("converting '%s' to %T", env, *val)
		}
		*vp = ival

	case *bool:
		bval, err := strconv.ParseBool(env)
		if err != nil {
			return fmt.Errorf("converting '%s' to %T", env, *val)
		}
		*vp = bval

	default:
		return fmt.Errorf("type not found: %T expected, got '%s'", *val, env)
	}
	return nil
}
