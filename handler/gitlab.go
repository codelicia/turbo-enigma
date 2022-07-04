package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"turboenigma/model"
	"turboenigma/provider"
)

type Gitlab struct {
	provider provider.Provider
}

func NewGitlab(provider provider.Provider) *Gitlab {
	return &Gitlab{provider: provider}
}

func (g *Gitlab) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
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

		if mr.EventType != "merge_request" {
			fmt.Fprint(writer, "We just care about merge_request events")
			return
		}

		if mr.ObjectAttributes.Action == "approved" {
			err = g.provider.NotifyMergeRequestApproved(mr)
			if err != nil {
				return err
			}

			fmt.Fprint(writer, "MR Approved OK....")

			return
		}

		// Filter events by "MergeRequest" opened
		if mr.ObjectAttributes.Action != "open" {
			fmt.Fprint(writer, "We can handle only open & approved events")
			return
		}

		err = g.provider.NotifyMergeRequestCreated(mr)
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

func jsonDecode(jsonString string) (mergeRequest model.MergeRequestInfo, err error) {
	err = json.Unmarshal([]byte(jsonString), &mergeRequest)

	return
}
