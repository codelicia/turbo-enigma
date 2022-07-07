package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"turboenigma/model"
)

type message struct {
	Text     string `json:"text"`
	IconURL  string `json:"icon_url,omitempty"`
	Username string `json:"username,omitempty"`
	Channel  string `json:"channel,omitempty"`
}

type search struct {
	Query string `json:"query"`
}

type LocatedMessage struct {
	channelID   string
	channelName string
	timestamp   string
	permalink   string
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Slack struct {
	client                                       HTTPClient
	notificationRules                            []model.NotificationRule
	reactionRules                                []model.ReactionRule
	webhookURL, message, avatar, username, token string
}

func NewSlack(client HTTPClient, notificationRules []model.NotificationRule, reactionRules []model.ReactionRule, webhookURL, token, message, avatar, username string) *Slack {
	return &Slack{
		client:            client,
		notificationRules: notificationRules,
		reactionRules:     reactionRules,
		webhookURL:        webhookURL,
		token:             token,
		message:           message,
		avatar:            avatar,
		username:          username,
	}
}

func (s *Slack) ReactToMessage(mergeRequest model.MergeRequestInfo, reactionRule model.ReactionRule) error {
	// Search for previous message
	var searchTerms = fmt.Sprintf("%s <%s|%s> by %s", s.message, mergeRequest.ObjectAttributes.URL, mergeRequest.ObjectAttributes.Title, mergeRequest.User.Name)

	locatedMessage, err := s.search(searchTerms)
	if err != nil {
		return err
	}

	s.postReaction(locatedMessage, reactionRule)

	return nil
}

func (s *Slack) NotifyMergeRequestOpened(mergeRequest model.MergeRequestInfo) error {
	channels := s.ChannelsForMergeRequest(mergeRequest)

	for _, channel := range channels {
		var m = message{
			Text:     fmt.Sprintf("%s <%s|%s> by %s", s.message, mergeRequest.ObjectAttributes.URL, mergeRequest.ObjectAttributes.Title, mergeRequest.User.Name),
			IconURL:  s.avatar,
			Username: s.username,
			Channel:  channel,
		}

		asJSON, err := json.Marshal(m)
		if err != nil {
			return err
		}

		err = s.sendMessage(asJSON)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Slack) ChannelsForMergeRequest(mergeRequest model.MergeRequestInfo) []string {
	channels := []string{}

	for _, config := range s.notificationRules {
		for _, mrLabel := range mergeRequest.Labels {
			for _, ruleLabel := range config.Labels {
				if mrLabel.Title == ruleLabel {
					channels = append(channels, config.Channel)
				}
			}
		}
	}

	return channels
}

func (s *Slack) NotifyMergeRequestMerged(mergeRequest model.MergeRequestInfo) error {
	return s.ReactToMessage(mergeRequest, model.ReactionRule{Action: "merge", Reaction: "white_tick"})
}

func (s *Slack) NotifyMergeRequestApproved(mergeRequest model.MergeRequestInfo) error {
	return s.ReactToMessage(mergeRequest, model.ReactionRule{Action: "approve", Reaction: "thumbsup"})
}

func (s *Slack) sendMessage(message []byte) error {
	req, err := http.NewRequest(http.MethodPost, s.webhookURL, bytes.NewBuffer(message))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response codeStatus code: %d, expected 200", resp.StatusCode)
	}

	return nil
}

func (s *Slack) search(searchTerms string) (message LocatedMessage, err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://slack.com/api/search.messages?query=%s", strings.ReplaceAll(searchTerms, " ", "%20")), bytes.NewBufferString(""))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	resp, err := s.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != http.StatusOK {
		return
	}

	searchResults, err := jsonDecodeMessage(string(b))

	return LocatedMessage{
		channelID:   searchResults.Messages.Matches[0].Channel.ID,
		channelName: searchResults.Messages.Matches[0].Channel.Name,
		timestamp:   searchResults.Messages.Matches[0].Ts,
		permalink:   searchResults.Messages.Matches[0].Permalink,
	}, nil
}

func (s *Slack) postReaction(message LocatedMessage, reactionRule model.ReactionRule) error {
	data := url.Values{}
	data.Set("channel", message.channelID)
	data.Set("name", reactionRule.Reaction)
	data.Set("timestamp", message.timestamp)

	req, err := http.NewRequest(http.MethodPost, "https://slack.com/api/reactions.add", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	resp, e := s.client.Do(req)
	if e != nil {
		return e
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response codeStatus code: %d, expected 200", resp.StatusCode)
	}

	return nil
}

func jsonDecodeMessage(jsonString string) (message model.SearchResult, err error) {
	err = json.Unmarshal([]byte(jsonString), &message)

	return
}
