package models

type NotificationConfig struct {
	Channel string `json:"channel"`
	Labels  []string `json:"labels"`
}
