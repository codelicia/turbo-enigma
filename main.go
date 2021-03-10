package main

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

//Checking that an environment variable is present or not.
func guardEnvVars() error {
	_, httpHost := os.LookupEnv("HTTP_HOST")
	if !httpHost {
		return errors.New("Missing HTTP_HOST in environment variable")
	}

	_, slackWebhookUrl := os.LookupEnv("SLACK_WEBHOOK_URL")
	if !slackWebhookUrl {
		return errors.New("Missing SLACK_WEBHOOK_URL in environment variable")
	}

	return nil
}

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

func defaultHandler(writer http.ResponseWriter, request *http.Request) {
	body, _ := ioutil.ReadAll(request.Body)
	fmt.Println("request Body:", string(body))

	var url = os.Getenv("SLACK_WEBHOOK_URL")
	var json = []byte(`{"text":"` + html.EscapeString(string(body)) + `"}`)

	err := postJson(url, json)
	if err != nil {
		fmt.Fprintf(writer, "Error: %s", err.Error())
		return
	}

	fmt.Fprintf(writer, "OK")
}

func main() {
	var server = fmt.Sprintf("0.0.0.0:%s", os.Getenv("HTTP_PORT"))

	fmt.Println("Server listening on", server)

	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(server, nil)
}
