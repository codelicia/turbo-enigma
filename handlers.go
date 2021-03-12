package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

func healthCheckOn(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "It is alive!")
}

func postOnSlack(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	assert(err)

	var url = os.Getenv("SLACK_WEBHOOK_URL")

	mr := jsonDecode(string(body))

	// Filter events by "MergeRequest" opened
	if mr.EventType != "merge_request" && mr.ObjectAttributes.Action != "open" {
		return
	}

	// Filtering by label
	var matchLabel = false
	for _, s := range mr.Labels {
		if s.Title == os.Getenv("MERGE_REQUEST_LABEL") {
			matchLabel = true
		}
	}

	if matchLabel == false {
		return
	}

	var message = []byte(`{"text": "` + os.Getenv("MESSAGE") + ` <` + html.EscapeString(mr.ObjectAttributes.URL) + `|` + html.EscapeString(mr.ObjectAttributes.Title) + `> by ` + html.EscapeString(mr.User.Name) + `"}`)

	err2 := postJson(url, message)

	if err2 != nil {
		fmt.Fprintf(writer, "Error -> %s", err.Error())
		return
	}

	fmt.Fprintf(writer, "OK")
}
