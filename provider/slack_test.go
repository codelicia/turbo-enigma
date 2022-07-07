package provider

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
	"turboenigma/model"

	"github.com/stretchr/testify/assert"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var (
	GetDoFunc func(req *http.Request) (*http.Response, error)
)

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestChannelsForMergeRequestSingleRule(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-just-testing.json")
	assert.Nil(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Nil(t, err)
	assert.Equal(t, "just-testing", mergeRequest.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"just-testing\"]}]"
	err = json.Unmarshal([]byte(jsonString), &notifications)
	assert.Nil(t, err)

	jsonString = "[{\"action\":\"approved\",\"reaction\":\"thumbsup\"}]"
	err = json.Unmarshal([]byte(jsonString), &reactions)
	assert.Nil(t, err)

	slack := NewSlack(
		http.DefaultClient,
		notifications,
		reactions,
		"https://testing.com",
		"New MR",
		"https://avatar",
		"Username",
		"Token",
	)

	assert.Equal(t, []string{"#tested"}, slack.ChannelsForMergeRequest(mergeRequest))
}

func TestChannelsForMergeRequestMultipleRules(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-enabling-team.json")
	assert.Nil(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Nil(t, err)
	assert.Equal(t, "Enabling Team", mergeRequest.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"just-testing\"]},{\"channel\":\"#multiple-rules\",\"labels\":[\"Enabling Team\"]}]"
	err = json.Unmarshal([]byte(jsonString), &notifications)
	assert.Nil(t, err)

	jsonString = "[{\"action\":\"approved\",\"reaction\":\"thumbsup\"}]"
	err = json.Unmarshal([]byte(jsonString), &reactions)
	assert.Nil(t, err)

	slack := NewSlack(
		http.DefaultClient,
		notifications,
		reactions,
		"https://testing.com",
		"New MR",
		"https://avatar",
		"Username",
		"Token",
	)

	assert.Equal(t, []string{"#multiple-rules"}, slack.ChannelsForMergeRequest(mergeRequest))
}

func TestChannelsForMergeRequestMultipleRulesWithMoreThanOneLabel(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-enabling-team.json")
	assert.Nil(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Nil(t, err)
	assert.Equal(t, "Enabling Team", mergeRequest.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"just-testing\", \"find-my-bug\"]},{\"channel\":\"#multiple-rules\",\"labels\":[\"Team Magie Lee\", \"Enabling Team\"]}]"
	err = json.Unmarshal([]byte(jsonString), &notifications)
	assert.Nil(t, err)

	jsonString = "[{\"action\":\"approved\",\"reaction\":\"thumbsup\"}]"
	err = json.Unmarshal([]byte(jsonString), &reactions)
	assert.Nil(t, err)

	slack := NewSlack(
		http.DefaultClient,
		notifications,
		reactions,
		"https://testing.com",
		"New MR",
		"https://avatar",
		"Username",
		"Token",
	)

	assert.Equal(t, []string{"#multiple-rules"}, slack.ChannelsForMergeRequest(mergeRequest))
}

func TestChannelsForMergeRequestNotMatchingLabel(t *testing.T) {
	var mergeRequest model.MergeRequestInfo
	var jsonString string
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	payload, err := ioutil.ReadFile("../payload/merge_request-open-just-testing.json")
	assert.Nil(t, err)

	err = json.Unmarshal(payload, &mergeRequest)
	assert.Nil(t, err)
	assert.Equal(t, "just-testing", mergeRequest.Labels[0].Title)

	jsonString = "[{\"channel\":\"#tested\",\"labels\":[\"something-else\"]}]"
	err = json.Unmarshal([]byte(jsonString), &notifications)
	assert.Nil(t, err)

	jsonString = "[{\"action\":\"approved\",\"reaction\":\"thumbsup\"}]"
	err = json.Unmarshal([]byte(jsonString), &reactions)
	assert.Nil(t, err)

	slack := NewSlack(
		http.DefaultClient,
		notifications,
		reactions,
		"https://testing.com",
		"New MR",
		"https://avatar",
		"Username",
		"Token",
	)

	assert.Equal(t, []string{}, slack.ChannelsForMergeRequest(mergeRequest))
}

func TestSearchForMessage(t *testing.T) {
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	json := `{"messages": {"matches": [{"channel": {"name": "channel-name", "id": "channel-id"}, "ts": "123", "permalink": "http://message-link"}]}}`
	client := &MockClient{}
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	slack := NewSlack(
		client,
		notifications,
		reactions,
		"https://testing.com",
		"New MR",
		"https://avatar",
		"Username",
		"Token",
	)

	locatedMessage, err := slack.search("New MergeRequest Created")
	assert.Nil(t, err)
	assert.Equal(t, "channel-name", locatedMessage.channelName)
	assert.Equal(t, "channel-id", locatedMessage.channelID)
	assert.Equal(t, "123", locatedMessage.timestamp)
	assert.Equal(t, "http://message-link", locatedMessage.permalink)
}

func TestPostReaction(t *testing.T) {
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	json := `{}`
	client := &MockClient{}
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	slack := NewSlack(client, notifications, reactions, "https://testing.com", "New MR", "https://avatar", "Username", "Token")

	locatedMessage := LocatedMessage{
		channelID:   "channel-id",
		channelName: "channel-name",
		timestamp:   "123",
		permalink:   "http://message-link",
	}

	err := slack.postReaction(locatedMessage, model.ReactionRule{Action: "approved", Reaction: "thumbsup"})
	assert.Nil(t, err)
}

func TestPostReactionFailsWithStatusCode(t *testing.T) {
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	json := `{}`
	client := &MockClient{}
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 404,
			Body:       r,
		}, nil
	}

	slack := NewSlack(client, notifications, reactions, "https://testing.com", "New MR", "https://avatar", "Username", "Token")

	locatedMessage := LocatedMessage{
		channelID:   "channel-id",
		channelName: "channel-name",
		timestamp:   "123",
		permalink:   "http://message-link",
	}

	err := slack.postReaction(locatedMessage, model.ReactionRule{Action: "approved", Reaction: "thumbsup"})
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "response codeStatus code: 404, expected 200")
}

func TestPostReactionFailsWithClientError(t *testing.T) {
	var notifications []model.NotificationRule
	var reactions []model.ReactionRule

	client := &MockClient{}
	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return nil, errors.New("Error from web server")
	}

	slack := NewSlack(client, notifications, reactions, "https://testing.com", "New MR", "https://avatar", "Username", "Token")

	locatedMessage := LocatedMessage{
		channelID:   "channel-id",
		channelName: "channel-name",
		timestamp:   "123",
		permalink:   "http://message-link",
	}

	err := slack.postReaction(locatedMessage, model.ReactionRule{Action: "approved", Reaction: "thumbsup"})
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "Error from web server")
}
