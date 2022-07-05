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

type Slack struct {
	client                                       *http.Client
	notificationRules                            []model.NotificationRule
	reactionRules                                []model.ReactionRule
	webhookURL, message, avatar, username, token string
}

func NewSlack(client *http.Client, notificationRules []model.NotificationRule, reactionRules []model.ReactionRule, webhookURL, token, message, avatar, username string) *Slack {
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
	// TODO: filter messages reaction per channel
	// channels := s.ChannelsForMergeRequest(mergeRequest)

	// Search for previous message
	var searchTerms = fmt.Sprintf("%s <%s|%s> by %s", s.message, mergeRequest.ObjectAttributes.URL, mergeRequest.ObjectAttributes.Title, mergeRequest.User.Name)

	// TODO: use reactionRule only when actually reacting
	s.search(searchTerms, reactionRule)

	return nil
}

func (s *Slack) GetReactionRules() []model.ReactionRule {
	return s.reactionRules
}

func (s *Slack) NotifyMergeRequestCreated(mergeRequest model.MergeRequestInfo) error {
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

func (s *Slack) search(searchTerms string, reactionRule model.ReactionRule) string {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://slack.com/api/search.messages?query=%s", strings.ReplaceAll(searchTerms, " ", "%20")), bytes.NewBufferString(""))
	if err != nil {
		return "a"
	}

	// TODO: make sure we have the correct token
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	resp, err := s.client.Do(req)
	if err != nil {
		return "b"
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "x"
	}

	// fmt.Println(string(b))

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintln(resp.StatusCode)
	}

	searchResults, err := jsonDecodeMessage(string(b))

	// TODO: filter per channel and get the latest message only
	var ts = searchResults.Messages.Matches[0].Ts
	var channelID = searchResults.Messages.Matches[0].Channel.ID

	fmt.Println(ts)
	fmt.Println(channelID)
	fmt.Println(searchResults.Messages.Matches[0].Channel.Name)
	fmt.Println(searchResults.Messages.Matches[0].Permalink)

	//---------------------------------------- Add reaction
	data := url.Values{}
	data.Set("channel", channelID)
	data.Set("name", reactionRule.Reaction)
	data.Set("timestamp", ts)

	req1, err1 := http.NewRequest(http.MethodPost, "https://slack.com/api/reactions.add", strings.NewReader(data.Encode()))
	if err1 != nil {
		return "a"
	}

	// TODO: make sure we have the correct token
	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req1.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))

	resp1, e := s.client.Do(req1)
	if e != nil {
		return "b"
	}
	defer resp1.Body.Close()

	if resp1.StatusCode != http.StatusOK {
		fmt.Printf("Reaction response codeStatus code: %d, expected 200", resp.StatusCode)
		return fmt.Sprintf("Reaction response codeStatus code: %d, expected 200", resp.StatusCode)
	}

	c, _ := io.ReadAll(resp.Body)

	fmt.Println(string(c))

	fmt.Printf("Reaction posted")
	return "end"
}

func jsonDecodeMessage(jsonString string) (message model.SearchResult, err error) {
	err = json.Unmarshal([]byte(jsonString), &message)

	return
}
