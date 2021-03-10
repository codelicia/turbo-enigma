package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func postJson(url string, json []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return err
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return nil
}

func main() {
	// Check for required environment variables
	guardEnvVars()

	var server = fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT"))

	fmt.Println("Server listening on", server)

	http.HandleFunc("/", postOnSlack)
	http.ListenAndServe(server, nil)
}
