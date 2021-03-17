package pkg_test

import (
	"io/ioutil"
	"testing"
	"turboenigma/pkg"
)

func TestJsonDecode(t *testing.T) {
	dat, err := ioutil.ReadFile("./payload/merge_request-open.json")
	pkg.Assert(err)

	mergeRequest := pkg.JsonDecode(string(dat))

	if mergeRequest.User.Username != "alexandre.eher" {
		t.Errorf("Expected ''; got %v", mergeRequest.User.Username)
	}
}
