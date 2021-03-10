package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

func postOnSlack(writer http.ResponseWriter, request *http.Request) {
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
