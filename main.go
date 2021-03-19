package main

import (
	"fmt"
	"net/http"
	"os"
	"turboenigma/pkg"
)

func init() {
	pkg.Client = &http.Client{}
}

func main() {
	if err := pkg.GuardEnvVars(); err != nil {
		panic(err)
	}

	var server = fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT"))

	fmt.Println("Server listening on", server)

	http.HandleFunc("/", pkg.PostOnSlack)
	http.HandleFunc("/healthcheck", pkg.HealthCheckOn)

	if err := http.ListenAndServe(server, nil); err != nil {
		panic(err)
	}
}
