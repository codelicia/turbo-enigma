package pkg_test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"turboenigma/pkg"
)

type SpyMessage struct {
	SendPullRequestEventFunc func(URL, title, author string) error
}

func (s *SpyMessage) SendPullRequestEvent(URL, title, author string) error {
	return s.SendPullRequestEventFunc(URL, title, author)
}

func TestPostOnSlack(t *testing.T) {
	dat, err := ioutil.ReadFile("./payload/merge_request-open.json")
	assert.Empty(t, err)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	os.Setenv("MERGE_REQUEST_LABEL", "just-testing")

	pkg.EnvManager, _ = pkg.NewEnv([]string{
		"MERGE_REQUEST_LABEL",
	})

	pkg.Message = &SpyMessage{
		SendPullRequestEventFunc: func(URL, title, author string) (err error) {
			assert.Equal(t, "https://gitlab.com/alexandre.eher/turbo-enigma/-/merge_requests/1", URL)
			assert.Equal(t, "Add LICENSE", title)
			assert.Equal(t, "Alexandre Eher", author)

			return
		},
	}
	pkg.PostOnSlack(rec, req)

	assert.Equal(t, "OK", rec.Body.String())
}

func TestPostOnSlackWithEmptyBody(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(""),
	)

	os.Setenv("MERGE_REQUEST_LABEL", "just-testing")

	pkg.EnvManager, _ = pkg.NewEnv([]string{
		"MERGE_REQUEST_LABEL",
	})

	pkg.Message = &SpyMessage{
		SendPullRequestEventFunc: func(URL, title, author string) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}
	pkg.PostOnSlack(rec, req)

	assert.Equal(t,"Error -> Body is missing\n", rec.Body.String())
	assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
}

func TestPostOnSlackWithIgnoredLabel(t *testing.T) {
	dat, err := ioutil.ReadFile("./payload/merge_request-open.json")
	assert.Empty(t, err)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	os.Setenv("MERGE_REQUEST_LABEL", "invalid-label")

	pkg.EnvManager, _ = pkg.NewEnv([]string{
		"MERGE_REQUEST_LABEL",
	})

	pkg.Message = &SpyMessage{
		SendPullRequestEventFunc: func(URL, title, author string) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}
	pkg.PostOnSlack(rec, req)

	assert.Equal(t,"We didn't find the right label", rec.Body.String())
}

func TestPostOnSlackWithNewIssue(t *testing.T) {
	dat, err := ioutil.ReadFile("./payload/issue-open.json")
	assert.Empty(t, err)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	os.Setenv("MERGE_REQUEST_LABEL", "just-testing")

	pkg.EnvManager, _ = pkg.NewEnv([]string{
		"MERGE_REQUEST_LABEL",
	})

	pkg.Message = &SpyMessage{
		SendPullRequestEventFunc: func(URL, title, author string) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}
	pkg.PostOnSlack(rec, req)

	assert.Equal(t,"We just care about new merge_requests", rec.Body.String())
}

func TestPostOnSlackWithMergeRequestApproved(t *testing.T) {
	dat, err := ioutil.ReadFile("./payload/merge_request-approved.json")
	assert.Empty(t, err)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader(string(dat)),
	)

	os.Setenv("MERGE_REQUEST_LABEL", "just-testing")

	pkg.EnvManager, _ = pkg.NewEnv([]string{
		"MERGE_REQUEST_LABEL",
	})

	pkg.Message = &SpyMessage{
		SendPullRequestEventFunc: func(URL, title, author string) (err error) {
			assert.FailNow(t, "Code should not reach this method")

			return
		},
	}
	pkg.PostOnSlack(rec, req)

	assert.Equal(t,"We just care about new merge_requests", rec.Body.String())
}