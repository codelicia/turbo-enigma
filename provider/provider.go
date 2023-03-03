// Package provider contains all the notification providers interface
// that can be used to send notifications to the slack client.
//
// The package contains a Provider interface that all the providers
// must implement.
package provider

import (
	"turboenigma/model"
)

type Provider interface {
	NotifyMergeRequestOpened(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestApproved(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestUnapproved(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestClose(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestReopen(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestUpdate(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestApproval(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestUnapproval(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestMerged(mergeRequest model.MergeRequestInfo) error
}
