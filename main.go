package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"turboenigma/models"
	"turboenigma/pkg"
	"turboenigma/pkg/message"
)

func main() {
	envManager, err := pkg.NewEnv([]string{
		"HTTP_PORT",
		"SLACK_WEBHOOK_URL",
		"MESSAGE",
		"NOTIFICATION_CONFIG",
		"SLACK_USERNAME",
		"SLACK_AVATAR_URL",
	})
	if err != nil {
		panic(err)
	}

	if (pkg.EnvManager.Get("MERGE_REQUEST_LABEL") != "") {
		fmt.Println("'MERGE_REQUEST_LABEL' is deprecated and will be removed soon.")
		fmt.Println("Please use 'NOTIFICATION_CONFIG' instead.")
	}

	var notifications []models.NotificationConfig
	err = json.Unmarshal([]byte(pkg.EnvManager.Get("NOTIFICATION_CONFIG")), &notifications)

	pkg.EnvManager = envManager
	pkg.Provider = message.NewSlack(
		http.DefaultClient,
		notifications,
		pkg.EnvManager.Get("SLACK_WEBHOOK_URL"),
		pkg.EnvManager.Get("MESSAGE"),
		pkg.EnvManager.Get("SLACK_AVATAR_URL"),
		pkg.EnvManager.Get("SLACK_USERNAME"),
	)

	var server = fmt.Sprintf("0.0.0.0:%s", pkg.EnvManager.Get("HTTP_PORT"))

	fmt.Println("Server listening on", server)

	http.HandleFunc("/", pkg.PostOnSlack)
	http.HandleFunc("/healthcheck", pkg.HealthCheckOn)

	if err := http.ListenAndServe(server, nil); err != nil {
		panic(err)
	}
}
