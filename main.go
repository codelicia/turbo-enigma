package main

import (
	"fmt"
	"net/http"
	"turboenigma/pkg"
	"turboenigma/pkg/message"
)

func main() {
	envManager, err := pkg.NewEnv([]string{
		"HTTP_PORT",
		"SLACK_WEBHOOK_URL",
		"MESSAGE",
		"MERGE_REQUEST_LABEL",
		"SLACK_USERNAME",
		"SLACK_AVATAR_URL",
	})
	if err != nil {
		panic(err)
	}

	pkg.EnvManager = envManager
	pkg.Message = message.NewSlack(
		http.DefaultClient,
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
