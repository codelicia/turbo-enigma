package main

import (
	"io/ioutil"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	dat, err := ioutil.ReadFile("./payload.json")
	assert(err)

	mergeRequest := jsonDecode(string(dat))

	if mergeRequest.User.Username != "alexandre.eher" {
		t.Errorf("Expected ''; got %v", mergeRequest.User.Username)
	}
}
