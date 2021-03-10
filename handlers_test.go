package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostOnSlack(t *testing.T) {
	req, err := http.NewRequest(
		"POST",
		"localhost:8080",
		strings.NewReader("{}"),
	)

	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rec := httptest.NewRecorder()

	postOnSlack(rec, req)
}
