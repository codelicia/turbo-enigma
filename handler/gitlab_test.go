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
	ReactToMessageFunc            func(mergeRequest model.MergeRequestInfo, reactionRule model.ReactionRule) error
	GetReactionRulesFunc          func() []model.ReactionRule
}

func (s *SpyProvider) NotifyMergeRequestCreated(mergeRequest model.MergeRequestInfo) error {
	return s.NotifyMergeRequestCreatedFunc(mergeRequest)
}

func (s *SpyProvider) ReactToMessage(mergeRequest model.MergeRequestInfo, reactionRule model.ReactionRule) error {
	return s.ReactToMessageFunc(mergeRequest, reactionRule)
}

func (s *SpyProvider) GetReactionRules() []model.ReactionRule {
	return []model.ReactionRule{{
		Action:   "approved",
		Reaction: "thumbsup",
	}}
}

func TestPostOnSlack(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-open-just-testing.json")
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

	assert.Equal(t, "We just care about merge_request events", recorder.Body.String())
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
		ReactToMessageFunc: func(mergeRequest model.MergeRequestInfo, reactionRule model.ReactionRule) (err error) {
			assert.Equal(t, reactionRule.Action, "approved")
			assert.Equal(t, reactionRule.Reaction, "thumbsup")

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "Reacting :thumbsup: to MR", recorder.Body.String())
}

func TestPostOnSlackWithMergeRequestRejected(t *testing.T) {
	dat, err := ioutil.ReadFile("../payload/merge_request-rejected.json")
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
		ReactToMessageFunc: func(mergeRequest model.MergeRequestInfo, reactionRule model.ReactionRule) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	assert.Equal(t, "We cannot handle rejected event action", recorder.Body.String())
}
