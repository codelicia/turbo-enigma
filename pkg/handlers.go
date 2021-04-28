package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"turboenigma/models"
	"turboenigma/pkg/message"
)

var(
	Provider message.Provider
)

func HealthCheckOn(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "It is alive!")
}

func PostOnSlack(writer http.ResponseWriter, request *http.Request) {
	err := func() (err error) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return err
		}

		if string(body) == "" {
			return fmt.Errorf("Body is missing")
		}

		mr, err := jsonDecode(string(body))
		if err != nil {
			return err
		}

		// Filter events by "MergeRequest" opened
		if mr.EventType != "merge_request" || mr.ObjectAttributes.Action != "open" {
			fmt.Fprint(writer, "We just care about new merge_requests")
			return
		}

		err = Provider.NotifyMergeRequestCreated(mr)
		if err != nil {
			return err
		}

		fmt.Fprint(writer, "OK")

		return
	}()

	if err != nil {
		http.Error(writer, fmt.Sprintf("Error -> %s", err.Error()), http.StatusBadRequest)
	}
}

func jsonDecode(jsonString string) (mergeRequest models.MergeRequestInfo, err error) {
	err = json.Unmarshal([]byte(jsonString), &mergeRequest)

	return
}
