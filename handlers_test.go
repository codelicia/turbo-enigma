package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func init() {
	Client = &MockClient{}
}

func TestPostOnSlack(t *testing.T) {
	var url = "http://turboenigma.com"

	GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(`ok`))),
		}, nil
	}

	dat, err := ioutil.ReadFile("./payload.json")
	assert(err)

	os.Setenv("SLACK_WEBHOOK_URL", url)
	os.Setenv("MERGE_REQUEST_LABEL", "just-testing")
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		strings.NewReader(string(dat)),
	)

	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rec := httptest.NewRecorder()

	postOnSlack(rec, req)

	if string(rec.Body.Bytes()) != "OK" {
		t.Errorf("rec.Body should be 'OK'; '%v' given", string(rec.Body.Bytes()))
	}
}

func TestPostOnSlackWithInvalidLabel(t *testing.T) {
	var url = "http://turboenigma.com"

	dat, err := ioutil.ReadFile("./payload.json")
	assert(err)

	os.Setenv("SLACK_WEBHOOK_URL", url)
	os.Setenv("MERGE_REQUEST_LABEL", "invalid-label")
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		strings.NewReader(string(dat)),
	)

	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rec := httptest.NewRecorder()

	postOnSlack(rec, req)

	if string(rec.Body.Bytes()) != "" {
		t.Errorf("rec.Body should be 'OK'; '%v' given", string(rec.Body.Bytes()))
	}
}
