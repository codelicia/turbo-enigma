package message

import (
	"turboenigma/models"
)

type Provider interface {
	// Deprecated
	SendPullRequestEvent(URL, title, author string) error
	NotifyMergeRequestCreated(mergeRequest models.MergeRequestInfo) error
}
