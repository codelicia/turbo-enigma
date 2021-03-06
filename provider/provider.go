package provider

import (
	"turboenigma/model"
)

type Provider interface {
	NotifyMergeRequestCreated(mergeRequest model.MergeRequestInfo) error
}
