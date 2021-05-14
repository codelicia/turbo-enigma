package main

import (
	"testing"
	"os"
)

func setupTestEnvironmentVariables() {
	os.Setenv("HTTP_PORT", "Just testing...")
	os.Setenv("MESSAGE", "Just testing...")
	os.Setenv("NOTIFICATION_RULES", "[{\"channel\":\"#test\", \"labels\":[\"test\"]}]")
	os.Setenv("SLACK_AVATAR_URL", "Just testing...")
	os.Setenv("SLACK_USERNAME", "Just testing...")
	os.Setenv("SLACK_WEBHOOK_URL", "Just testing...")
}

func TestEnvironmentMissingHttpPort(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("HTTP_PORT")

	InitEnvironment()

	t.Error("should panic about missing env: HTTP_PORT")
}

func TestEnvironmentMissingMessage(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("MESSAGE")

	InitEnvironment()

	t.Error("should panic about missing env: MESSAGE")
}


func TestEnvironmentMissingNotificationRules(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("NOTIFICATION_RULES")

	InitEnvironment()

	t.Error("should panic about missing env: NOTIFICATION_RULES")
}


func TestEnvironmentMissingSlackAvatarUrl(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("SLACK_AVATAR_URL")

	InitEnvironment()

	t.Error("should panic about missing env: SLACK_AVATAR_URL")
}


func TestEnvironmentMissingSlackUsername(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("SLACK_USERNAME")

	InitEnvironment()

	t.Error("should panic about missing env: SLACK_USERNAME")
}


func TestEnvironmentMissingSlackWebhookUrl(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("SLACK_WEBHOOK_URL")

	InitEnvironment()

	t.Error("should panic about missing env: SLACK_WEBHOOK_URL")
}
