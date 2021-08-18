package handler_test

import (
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
	NotifyMergeRequestCreatedFunc func(mergeRequest model.MergeRequestInfo) error
}

func (s *SpyProvider) NotifyMergeRequestCreated(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestCreatedFunc(mergeRequest)
}

func TestPostOnSlack(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-open.json")
	assert.Empty(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestCreatedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
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
		NotifyMergeRequestCreatedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
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
	assert.Empty(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestCreatedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "We just care about new merge_requests", recorder.Body.String())
}

func TestPostOnSlackWithMergeRequestApproved(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-approved.json")
	assert.Empty(t, err)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	provider := &SpyProvider{
		NotifyMergeRequestCreatedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "We just care about new merge_requests", recorder.Body.String())
}
