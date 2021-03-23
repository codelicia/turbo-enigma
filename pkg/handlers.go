package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"turboenigma/pkg/message"
)

var(
	Message message.Message
)

func HealthCheckOn(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "It is alive!")
}

func PostOnSlack(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Error -> %s", err.Error()), http.StatusBadRequest)
	}

	if string(body) == "" {
		http.Error(writer, "Body is missing", http.StatusBadRequest)
		return
	}

	mr := JSONDecode(string(body))

	// Filter events by "MergeRequest" opened
	if mr.EventType != "merge_request" || mr.ObjectAttributes.Action != "open" {
		fmt.Fprint(writer, "We just care about new merge_requests")
		return
	}

	// Filtering by label
	var matchLabel = false
	for _, s := range mr.Labels {
		if s.Title == os.Getenv("MERGE_REQUEST_LABEL") {
			matchLabel = true
			break
		}
	}

	if !matchLabel {
		fmt.Fprint(writer, "We didn't find the right label")
		return
	}

	err = Message.SendPullRequestEvent(mr.ObjectAttributes.URL, mr.ObjectAttributes.Title, mr.User.Name)
	if err != nil {
		http.Error(writer, fmt.Sprintf("Error -> %s", err.Error()), http.StatusBadRequest)
	}

	fmt.Fprint(writer, "OK")
}
