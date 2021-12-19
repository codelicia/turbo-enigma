package model

type NotificationRule struct {
	Channel string   `json:"channel"`
	Labels  []string `json:"labels"`
}
