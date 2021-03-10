package main

import (
	"os"
	"testing"
)

func TestEnvironment(t *testing.T) {
	err := guardEnvVars()
	if err == nil {
		t.Error("It fail to recognize that 'HTTP_PORT' is missing")
	}
	if err.Error() != "Missing HTTP_HOST in environment variable" {
		t.Error("HTTP_HOST should be missing, but it is not.")
	}

	os.Setenv("HTTP_HOST", "8980")

	err2 := guardEnvVars()
	if err2 == nil {
		t.Error("It fail to recognize that 'SLACK_WEBHOOK_URL' is missing")
	}
	if err2.Error() != "Missing SLACK_WEBHOOK_URL in environment variable" {
		t.Error("SLACK_WEBHOOK_URL should be missing, but it is not.")
	}

	os.Setenv("SLACK_WEBHOOK_URL", "http://turboenigma.localhost")

	err3 := guardEnvVars()
	if err3 != nil {
		t.Error("Environment variables was expected to be OK.")
	}
}
