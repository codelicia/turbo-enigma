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
	"turboenigma/provider"

	"github.com/stretchr/testify/assert"
)

// todo move SpyProvider to another file
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

func usePayload(t *testing.T, filepath string) string {
	t.Helper()
	content, err := ioutil.ReadFile(filepath)
	assert.Nil(t, err)

	return string(content)
}

func doRequest(provider provider.Provider, content string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://some-url.com", strings.NewReader(content))

	handler.NewGitlab(provider).ServeHTTP(recorder, request)

	return recorder
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
	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.Equal(t, "https://gitlab.com/alexandre.eher/turbo-enigma/-/merge_requests/1", mergeRequest.ObjectAttributes.URL)
			assert.Equal(t, "Add LICENSE", mergeRequest.ObjectAttributes.Title)
			assert.Equal(t, "Alexandre Eher", mergeRequest.User.Name)

			return
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/merge_request-open-just-testing.json"))

	assert.Equal(t, "OK", recorder.Body.String())
}

func TestPostOnSlackWithEmptyBody(t *testing.T) {
	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	recorder := doRequest(provider, "")

	assert.Equal(t, "Error -> Body is missing\n", recorder.Body.String())
	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
}

func TestPostOnSlackWithNewIssue(t *testing.T) {
	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/issue-open.json"))

	assert.Equal(t, "We just care about merge_request events", recorder.Body.String())
}

func TestPostOnSlackWithMergeRequestMerged(t *testing.T) {
	provider := &SpyProvider{
		NotifyMergeRequestMergedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.Equal(t, mergeRequest.ObjectAttributes.Action, "merge")
			return
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/merge_request-merge.json"))

	assert.Equal(t, "Reacting to merge event", recorder.Body.String())
}

func TestApprovedAction(t *testing.T) {

	t.Run("Happy path", func(t *testing.T) {
		provider := &SpyProvider{
			NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
				assert.Equal(t, mergeRequest.ObjectAttributes.Action, "approved")
				return
			},
		}

		recorder := doRequest(provider, usePayload(t, "../payload/merge_request-approved.json"))

		assert.Equal(t, "Reacting to approved event", recorder.Body.String())
	})

	t.Run("Failed", func(t *testing.T) {
		provider := &SpyProvider{
			NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
				return errors.New("NotifyMergeRequestApproved failed (on purpose)")
			},
		}

		recorder := doRequest(provider, usePayload(t, "../payload/merge_request-approved.json"))

		assert.Equal(t, "Error -> NotifyMergeRequestApproved failed (on purpose)\n", recorder.Body.String())
	})

}

// TODO(malukenho): get real payload for these events, right now I'm changing just the ObjectAttributes.Action
// func TestPostOnSlackWithMergeRequestUnapproved(t *testing.T) {
// 	provider := &SpyProvider{
// 		NotifyMergeRequestUnapprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
// 			assert.Equal(t, mergeRequest.ObjectAttributes.Action, "unapproved")
// 			return
// 		},
// 	}

// 	recorder := doRequest(provider, usePayload(t, "../payload/merge_request-approved.json"))

// 	assert.Equal(t, "Reacting to unapproved event", recorder.Body.String())
// }

func TestPostOnSlackWithMergeRequestRejected(t *testing.T) {
	provider := &SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")
			return
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/merge_request-rejected.json"))

	assert.Equal(t, "We cannot handle rejected event action", recorder.Body.String())
}

func TestPostOnSlackWithReactToMessageFailure(t *testing.T) {
	provider := &SpyProvider{
		NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			return errors.New("Error from ReactToMessage")
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/merge_request-rejected.json"))

	assert.Equal(t, "We cannot handle rejected event action", recorder.Body.String())
}
