package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"turboenigma/handler"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"http://some-url.com",
		strings.NewReader("test"),
	)

	handler.NewHealthCheck().ServeHTTP(recorder, request)

	assert.Equal(t, "It is alive!", recorder.Body.String())
}
