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

	if string(body) == "" {
		http.Error(writer, "Body is missing", http.StatusBadRequest)
		return
	}

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

	var template = "{'text': '%s <%s|%s> by %s', 'icon_url': '%s', 'username': '%s'}"

	var formating = fmt.Sprintf(
		template,
		os.Getenv("MESSAGE"),
		html.EscapeString(mr.ObjectAttributes.URL),
		html.EscapeString(mr.ObjectAttributes.Title),
		html.EscapeString(mr.User.Name),
		"https://avatars.githubusercontent.com/u/46966179?s=200&v=4",
		"codelicia/turbo-enigma",
	)
	var message = []byte(formating)

	err2 := postJson(url, message)

	if err2 != nil {
		http.Error(writer, fmt.Sprintf("Error -> %s", err2.Error()), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "OK")
}
