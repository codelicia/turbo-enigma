package provider

import (
	"turboenigma/model"
)

type Provider interface {
	// Deprecated: please use NotifyMergeRequestCreated instead
	SendPullRequestEvent(URL, title, author string) error
	NotifyMergeRequestCreated(mergeRequest model.MergeRequestInfo) error
}
