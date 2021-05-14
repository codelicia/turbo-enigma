package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"turboenigma/handler"
	"turboenigma/model"
	"turboenigma/provider"
)

func InitEnvironment() *Env {
	env, err := NewEnv([]string{
		"HTTP_PORT",
		"MESSAGE",
		"NOTIFICATION_RULES",
		"SLACK_AVATAR_URL",
		"SLACK_USERNAME",
		"SLACK_WEBHOOK_URL",
	})
	if err != nil {
		panic(err)
	}

	return env
}


func main() {
	EnvManager = InitEnvironment()

	var notificationRules []model.NotificationRule
	err := json.Unmarshal([]byte(EnvManager.Get("NOTIFICATION_RULES")), &notificationRules)
	if err != nil {
		panic(err)
	}

	slack := provider.NewSlack(
		http.DefaultClient,
		notificationRules,
		EnvManager.Get("SLACK_WEBHOOK_URL"),
		EnvManager.Get("MESSAGE"),
		EnvManager.Get("SLACK_AVATAR_URL"),
		EnvManager.Get("SLACK_USERNAME"),
	)

	http.HandleFunc("/", handler.NewGitlab(slack).ServeHTTP)
	http.HandleFunc("/healthcheck", handler.NewHealthCheck().ServeHTTP)

	address := fmt.Sprintf("0.0.0.0:%s", EnvManager.Get("HTTP_PORT"))
	fmt.Println("Server listening on", address)

	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
