package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"turboenigma/model"

	"github.com/stretchr/testify/assert"
)

func TestChannelsForMergeRequest(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open.json")
	assert.Empty(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Empty(t, err)
	assert.Equal(t, "just-testing", mergeRequest.Labels[0].Title)

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

	assert.Equal(t, []string{"#tested"}, slack.ChannelsForMergeRequest(mergeRequest))
}

func TestChannelsForMergeRequestNotMatchingLabel(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open.json")
	assert.Empty(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Empty(t, err)
	assert.Equal(t, "just-testing", mergeRequest.Labels[0].Title)

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

	assert.Equal(t, []string{}, slack.ChannelsForMergeRequest(mergeRequest))
}
