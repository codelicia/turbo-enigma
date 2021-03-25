package message

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type message struct {
	Text string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
	Username string `json:"username,omitempty"`
	Channel string `json:"channel,omitempty"`
}

type Slack struct {
	client *http.Client
	webhookURL, message, avatar, username string
}

func NewSlack(client *http.Client, webhookURL, message, avatar, username string) *Slack {
	return &Slack{
		client:     client,
		webhookURL: webhookURL,
		message:    message,
		avatar:     avatar,
		username:   username,
	}
}

func (s *Slack) SendPullRequestEvent(URL, title, author string) error {
	var m = message{
		Text: fmt.Sprintf("%s <%s|%s> by %s", s.message, URL, title, author),
		IconURL: s.avatar,
		Username: s.username,
	}

	asJSON, err := json.Marshal(m)
	if err != nil {
		return err
	}

	return s.sendMessage(asJSON)
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

