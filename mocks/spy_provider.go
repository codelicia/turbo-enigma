package mocks

import (
	"turboenigma/model"
)

type SpyProvider struct {
	NotifyMergeRequestOpenedFunc     func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestApprovedFunc   func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestUnapprovedFunc func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestCloseFunc      func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestReopenFunc     func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestUpdateFunc     func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestApprovalFunc   func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestUnapprovalFunc func(mergeRequest model.MergeRequestInfo) error
	NotifyMergeRequestMergedFunc     func(mergeRequest model.MergeRequestInfo) error
}

func (s *SpyProvider) NotifyMergeRequestOpened(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestOpenedFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestApproved(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestApprovedFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestUnapproved(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestUnapprovalFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestClose(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestCloseFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestReopen(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestReopenFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestUpdate(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestUpdateFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestApproval(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestApprovalFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestUnapproval(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestUnapprovalFunc(mergeRequest)
}
func (s *SpyProvider) NotifyMergeRequestMerged(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestMergedFunc(mergeRequest)
}
