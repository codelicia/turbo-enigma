package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEnvironmentVariable(t *testing.T) {
	tt := []struct {
		envVar   string
		envValue string
	}{
		{"HTTP_PORT", "8980"},
		{"SLACK_WEBHOOK_URL", "http://turboenigma.localhost"},
		{"MESSAGE", "I need review on "},
		{"MERGE_REQUEST_LABEL", "codelicia-team"},
		{"SLACK_USERNAME", "codelicia/turbo-enigma"},
		{"SLACK_AVATAR_URL", "https://avatars.githubusercontent.com/u/46966179?s=200&v=4"},
	}

	for _, tc := range tt {
		t.Run(tc.envVar, func(t *testing.T) {
			err := guardEnvVars()
			if err == nil {
				t.Errorf("It fail to recognize that '%s' is missing", tc.envVar)
			}
			if err.Error() != fmt.Sprintf("Missing %s in environment variable", tc.envVar) {
				t.Errorf("%s should be missing, but it is not.", tc.envVar)
			}

			os.Setenv(tc.envVar, tc.envValue)
		})
	}

	if guardEnvVars() != nil {
		t.Error("Environment variables was expected to be OK.")
	}
}
