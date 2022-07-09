package handler_test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"turboenigma/handler"
	"turboenigma/mocks"
	"turboenigma/model"
	"turboenigma/provider"

	"github.com/stretchr/testify/assert"
)

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

// Describe all cases, except "open"
var actions = [8]string{
	"close",
	"reopen",
	"update",
	"approved",
	"unapproved",
	"approval",
	"unapproval",
	"merge",
}

func TestPostOnSlack(t *testing.T) {
	provider := &mocks.SpyProvider{
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
	provider := &mocks.SpyProvider{
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
	provider := &mocks.SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/issue-open.json"))

	assert.Equal(t, "We just care about merge_request events", recorder.Body.String())
}

func TestPostOnSlackWithMergeRequestMerged(t *testing.T) {
	provider := &mocks.SpyProvider{
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
		provider := &mocks.SpyProvider{
			NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
				assert.Equal(t, mergeRequest.ObjectAttributes.Action, "approved")
				return
			},
		}

		recorder := doRequest(provider, usePayload(t, "../payload/merge_request-approved.json"))

		assert.Equal(t, "Reacting to approved event", recorder.Body.String())
	})

	t.Run("Failed", func(t *testing.T) {
		provider := &mocks.SpyProvider{
			NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
				return errors.New("NotifyMergeRequestApproved failed (on purpose)")
			},
		}

		recorder := doRequest(provider, usePayload(t, "../payload/merge_request-approved.json"))

		assert.Equal(t, "Error -> NotifyMergeRequestApproved failed (on purpose)\n", recorder.Body.String())
	})

}

func TestGenericAction(t *testing.T) {

	configureForAction := func(t *testing.T, action string) func(mergeRequest model.MergeRequestInfo) (err error) {
		return func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.Equal(t, mergeRequest.ObjectAttributes.Action, action)
			return
		}
	}
	for _, action := range actions {

		t.Run(fmt.Sprintf("Happy path %s", action), func(t *testing.T) {
			provider := &mocks.SpyProvider{
				NotifyMergeRequestApprovalFunc:   configureForAction(t, action),
				NotifyMergeRequestApprovedFunc:   configureForAction(t, action),
				NotifyMergeRequestCloseFunc:      configureForAction(t, action),
				NotifyMergeRequestMergedFunc:     configureForAction(t, action),
				NotifyMergeRequestOpenedFunc:     configureForAction(t, action),
				NotifyMergeRequestReopenFunc:     configureForAction(t, action),
				NotifyMergeRequestUnapprovalFunc: configureForAction(t, action),
				NotifyMergeRequestUnapprovedFunc: configureForAction(t, action),
				NotifyMergeRequestUpdateFunc:     configureForAction(t, action),
			}
			recorder := doRequest(provider, usePayload(t, fmt.Sprintf("../payload/merge_request-%s.json", action)))

			assert.Equal(t, fmt.Sprintf("Reacting to %s event", action), recorder.Body.String())
		})

	}

	configureActionForException := func(t *testing.T, action string) func(mergeRequest model.MergeRequestInfo) (err error) {
		return func(mergeRequest model.MergeRequestInfo) (err error) {
			return errors.New(fmt.Sprintf("NotifyMergeRequest%s failed (on purpose)", action))
		}
	}
	for _, action := range actions {

		t.Run("Failed", func(t *testing.T) {
			provider := &mocks.SpyProvider{
				NotifyMergeRequestApprovalFunc:   configureActionForException(t, action),
				NotifyMergeRequestApprovedFunc:   configureActionForException(t, action),
				NotifyMergeRequestCloseFunc:      configureActionForException(t, action),
				NotifyMergeRequestMergedFunc:     configureActionForException(t, action),
				NotifyMergeRequestOpenedFunc:     configureActionForException(t, action),
				NotifyMergeRequestReopenFunc:     configureActionForException(t, action),
				NotifyMergeRequestUnapprovalFunc: configureActionForException(t, action),
				NotifyMergeRequestUnapprovedFunc: configureActionForException(t, action),
				NotifyMergeRequestUpdateFunc:     configureActionForException(t, action),
			}

			recorder := doRequest(provider, usePayload(t, fmt.Sprintf("../payload/merge_request-%s.json", action)))

			assert.Equal(t, fmt.Sprintf("Error -> NotifyMergeRequest%s failed (on purpose)\n", action), recorder.Body.String())
		})

	}

}

func TestPostOnSlackWithMergeRequestRejected(t *testing.T) {
	provider := &mocks.SpyProvider{
		NotifyMergeRequestOpenedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			assert.FailNow(t, "Code should not reach this method")
			return
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/merge_request-rejected.json"))

	assert.Equal(t, "We cannot handle rejected event action", recorder.Body.String())
}

func TestPostOnSlackWithReactToMessageFailure(t *testing.T) {
	provider := &mocks.SpyProvider{
		NotifyMergeRequestApprovedFunc: func(mergeRequest model.MergeRequestInfo) (err error) {
			return errors.New("Error from ReactToMessage")
		},
	}

	recorder := doRequest(provider, usePayload(t, "../payload/merge_request-rejected.json"))

	assert.Equal(t, "We cannot handle rejected event action", recorder.Body.String())
}
