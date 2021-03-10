package main

import (
	"errors"
	"os"
)

//Checking that an environment variable is present or not.
func guardEnvVars() error {
	_, httpHost := os.LookupEnv("HTTP_PORT")
	if !httpHost {
		return errors.New("Missing HTTP_PORT in environment variable")
	}

	_, slackWebhookUrl := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !slackWebhookUrl {
		return errors.New("Missing SLACK_WEBHOOK_URL in environment variable")
	}

	return nil
}
