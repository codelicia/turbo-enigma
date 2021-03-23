package pkg

import (
	"fmt"
	"os"
)

var (
	EnvManager *Env
)

type Env struct {
	envs map[string]string
}

func NewEnv(mandatoryEnvs []string) (*Env, error) {
	envs, err := mapEnvs(mandatoryEnvs)
	if err != nil {
		return nil, err
	}

	return &Env{envs: envs}, nil
}

func mapEnvs(envs []string) (envMap map[string]string, err error) {
	envMap = make(map[string]string)
	for _, env := range envs {
		value, ok := os.LookupEnv(env)
		if !ok {
			return envMap, fmt.Errorf("missing %s in environment variable", env)
		}

		envMap[env] = value
	}

	return
}

func (e *Env) Get(key string) (value string) {
	value, ok := e.envs[key]
	if !ok {
		panic(fmt.Errorf("missing %s in environment variable", key))
	}

	return
}
