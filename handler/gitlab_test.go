package handler_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"turboenigma/handler"
	"turboenigma/model"

	"github.com/stretchr/testify/assert"
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

func TestPostOnSlack(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-open-just-testing.json")
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.Equal(t, "https://gitlab.com/alexandre.eher/turbo-enigma/-/merge_requests/1", mergeRequest.ObjectAttributes.URL)
			assert.Equal(t, "Add LICENSE", mergeRequest.ObjectAttributes.Title)
			assert.Equal(t, "Alexandre Eher", mergeRequest.User.Name)

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "OK", recorder.Body.String())
}

func TestPostOnSlackWithEmptyBody(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(""),
	)

	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "Error -> Body is missing\n", recorder.Body.String())
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestPostOnSlackWithNewIssue(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/issue-open.json")
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "We just care about merge_request events", recorder.Body.String())
}

func TestPostOnSlackWithMergeRequestMerged(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-merge.json")
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
		NotifyMergeRequestMergedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.Equal(t, mergeRequest.ObjectAttributes.Action, "merge")
			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "Reacting to merge event", recorder.Body.String())
}

func TestPostOnSlackWithMergeRequestApproved(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-approved.json")
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
		NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.Equal(t, mergeRequest.ObjectAttributes.Action, "approved")
			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "Reacting to approved event", recorder.Body.String())
}

func TestPostOnSlackWithMergeRequestRejected(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-rejected.json")
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
		NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "We cannot handle rejected event action", recorder.Body.String())
}

func TestPostOnSlackWithReactToMessageFailure(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-rejected.json")
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
		NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			return errors.New("Error from ReactToMessage")
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "We cannot handle rejected event action", recorder.Body.String())
}
