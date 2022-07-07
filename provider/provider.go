package provider

import (
	"turboenigma/model"
)

type Provider interface {
	NotifyMergeRequestOpened(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestMerged(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestApproved(mergeRequest model.MergeRequestInfo) error
}
