package pkg

import (
	"encoding/json"
	"turboenigma/models"
)

func JSONDecode(jsonString string) (mergeRequest models.MergeRequestInfo) {
	err := json.Unmarshal([]byte(jsonString), &mergeRequest)
	Assert(err)

	return
}
