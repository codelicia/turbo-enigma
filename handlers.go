package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

func postOnSlack(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	assert(err)

	var url = os.Getenv("SLACK_WEBHOOK_URL")

	mr := jsonDecode(string(body))

	// Filtering by label
	for _, s := range mr.Labels {
		if s.Title != "just-testing" {
			return
		}
	}

	var message = []byte(`{"text": "Merge Request Created by ` + html.EscapeString(mr.User.Name) + `"}`)

	err2 := postJson(url, message)

	if err2 != nil {
		fmt.Fprintf(writer, "Error -> %s", err.Error())
		return
	}

	fmt.Fprintf(writer, "OK")
}
