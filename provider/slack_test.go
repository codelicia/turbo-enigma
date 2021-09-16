package provider

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"turboenigma/model"

	"github.com/stretchr/testify/assert"
)

func TestChannelsForMergeRequestSingleRule(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-just-testing.json")
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

func TestChannelsForMergeRequestMultipleRules(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-enabling-team.json")
	assert.Empty(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Empty(t, err)
	assert.Equal(t, "Enabling Team", mergeRequest.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"just-testing\"]},{\"channel\":\"#multiple-rules\",\"labels\":[\"Enabling Team\"]}]"
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

	assert.Equal(t, []string{"#multiple-rules"}, slack.ChannelsForMergeRequest(mergeRequest))
}

func TestChannelsForMergeRequestMultipleRulesWithMoreThanOneLabel(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-enabling-team.json")
	assert.Empty(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Empty(t, err)
	assert.Equal(t, "Enabling Team", mergeRequest.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"just-testing\", \"find-my-bug\"]},{\"channel\":\"#multiple-rules\",\"labels\":[\"Team Magie Lee\", \"Enabling Team\"]}]"
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

	assert.Equal(t, []string{"#multiple-rules"}, slack.ChannelsForMergeRequest(mergeRequest))
}

func TestChannelsForMergeRequestNotMatchingLabel(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-just-testing.json")
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
