package main

import (
	"encoding/json"
)

func jsonDecode(jsonString string) MergeRequestInfo {
	var mergeRequest MergeRequestInfo

	if err := json.Unmarshal([]byte(jsonString), &mergeRequest); err != nil {
		panic(err)
	}

	return mergeRequest
}
