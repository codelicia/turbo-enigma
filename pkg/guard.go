package pkg

import (
	"fmt"
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
		if _, value := os.LookupEnv(env); !value {
			return fmt.Errorf("missing %s in environment variable", env)
		}
	}

	return nil
}

func Assert(e error) {
	if e != nil {
		panic(e)
	}
}
