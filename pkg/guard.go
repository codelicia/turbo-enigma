package pkg

import (
	"errors"
	"os"
)

func GuardEnvVars() error {
	envs := []string{
		"HTTP_PORT",
		"SLACK_WEBHOOK_URL",
		"MESSAGE",
		"MERGE_REQUEST_LABEL",
		"SLACK_USERNAME",
		"SLACK_AVATAR_URL",
	}

	for _, env := range envs {
		_, value := os.LookupEnv(env)
		if !value {
			return errors.New("Missing " + env + " in environment variable")
		}
	}

	return nil
}

func Assert(e error) {
	if e != nil {
		panic(e)
	}
}
