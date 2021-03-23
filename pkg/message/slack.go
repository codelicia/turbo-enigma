package message

import (
	"bytes"
	"fmt"
	"html"
	"net/http"
)

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
	var template = "{'text': '%s <%s|%s> by %s', 'icon_url': '%s', 'username': '%s'}"

	var message = fmt.Sprintf(
		template,
		s.message,
		html.EscapeString(URL),
		html.EscapeString(title),
		html.EscapeString(author),
		s.avatar,
		s.username,
	)

	return s.sendMessage(message)
}

func (s *Slack) sendMessage(message string) error {
	req, err := http.NewRequest(http.MethodPost, s.webhookURL, bytes.NewBuffer([]byte(message)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response codeStatus code: %d, expected 200", resp.StatusCode)
	}

	return nil
}

