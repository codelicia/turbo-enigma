package pkg

import (
	"encoding/json"
)

func JsonDecode(jsonString string) MergeRequestInfo {
	var mergeRequest MergeRequestInfo

	if err := json.Unmarshal([]byte(jsonString), &mergeRequest); err != nil {
		panic(err)
	}

	return mergeRequest
}
