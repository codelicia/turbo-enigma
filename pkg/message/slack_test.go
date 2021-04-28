package message

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"turboenigma/models"

	"github.com/stretchr/testify/assert"
)

func TestChannelsForMergeRequest(t *testing.T) {
	var merge_request models.MergeRequestInfo
	var jsonString string
	var notifications []models.NotificationConfig

	payload, err := ioutil.ReadFile("../payload/merge_request-open.json")
	assert.Empty(t, err)

	err = json.Unmarshal(payload, &merge_request)
	assert.Empty(t, err)
	assert.Equal(t, "just-testing", merge_request.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"just-testing\"]}]"
	err = json.Unmarshal([]byte(jsonString), &notifications)
	assert.Empty(t, err)

	slack := NewSlack(
		http.DefaultClient,
		notifications,
		"https://testing.com",
		"New MR",
		"https://avatar",
		"Username",
	)

	assert.Equal(t, []string{"#tested"}, slack.ChannelsForMergeRequest(merge_request))
}

func TestChannelsForMergeRequestNotMatchingLabel(t *testing.T) {
	var merge_request models.MergeRequestInfo
	var jsonString string
	var notifications []models.NotificationConfig

	payload, err := ioutil.ReadFile("../payload/merge_request-open.json")
	assert.Empty(t, err)

	err = json.Unmarshal(payload, &merge_request)
	assert.Empty(t, err)
	assert.Equal(t, "just-testing", merge_request.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"something-else\"]}]"
	err = json.Unmarshal([]byte(jsonString), &notifications)
	assert.Empty(t, err)

	slack := NewSlack(
		http.DefaultClient,
		notifications,
		"https://testing.com",
		"New MR",
		"https://avatar",
		"Username",
	)

	assert.Equal(t, []string{}, slack.ChannelsForMergeRequest(merge_request))
}
