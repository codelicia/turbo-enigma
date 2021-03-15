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
	pkg.GuardEnvVars()

	var server = fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT"))

	fmt.Println("Server listening on", server)

	http.HandleFunc("/", pkg.PostOnSlack)
	http.HandleFunc("/healthcheck", pkg.HealthCheckOn)
	http.ListenAndServe(server, nil)
}
