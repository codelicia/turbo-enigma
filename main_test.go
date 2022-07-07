package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func setupTestEnvironmentVariables() {
	os.Setenv("HTTP_PORT", "Just testing...")
	os.Setenv("MESSAGE", "Just testing...")
	os.Setenv("NOTIFICATION_RULES", `[{"channel":"#test", "labels":["test"]}]`)
	os.Setenv("REACTION_RULES", `[{"action":"approved", "reaction":"thumbsup"}]`)
	os.Setenv("SLACK_AVATAR_URL", "Just testing...")
	os.Setenv("SLACK_USERNAME", "Just testing...")
	os.Setenv("SLACK_WEBHOOK_URL", "Just testing...")
	os.Setenv("SLACK_TOKEN", "xoxp-slack-token")
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

func TestEnvironmentMissingReactionRules(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("REACTION_RULES")

	InitEnvironment()

	t.Error("should panic about missing env: REACTION_RULES")
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

func TestEnvironmentMissingSlackToken(t *testing.T) {
	defer func() { recover() }()

	setupTestEnvironmentVariables()
	os.Unsetenv("SLACK_TOKEN")

	InitEnvironment()

	t.Error("should panic about missing env: SLACK_TOKEN")
}

func TestEnvironmentWithInvalidReactionRulesJson(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			ok := r.(*json.SyntaxError)
			assert.ErrorContains(t, ok, "unexpected end of JSON input")
		}
	}()

	setupTestEnvironmentVariables()
	os.Unsetenv("REACTION_RULES")
	os.Setenv("REACTION_RULES", `{"wrong`)

	main()

	t.Error("should panic invalid json on REACTION_RULES")
}

func TestEnvironmentWithInvalidNotificationRuleJson(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			ok := r.(*json.SyntaxError)
			assert.ErrorContains(t, ok, "unexpected end of JSON input")
		}
	}()

	setupTestEnvironmentVariables()
	os.Unsetenv("NOTIFICATION_RULES")
	os.Setenv("NOTIFICATION_RULES", `{"wrong`)

	main()

	t.Error("should panic invalid json on NOTIFICATION_RULES")
}
