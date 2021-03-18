package pkg

import (
	"encoding/json"
	"turboenigma/models"
)

func JsonDecode(jsonString string) models.MergeRequestInfo {
	var mergeRequest models.MergeRequestInfo

	if err := json.Unmarshal([]byte(jsonString), &mergeRequest); err != nil {
		panic(err)
	}

	return mergeRequest
}
