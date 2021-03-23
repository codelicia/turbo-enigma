package pkg_test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"turboenigma/pkg"
)

func TestJsonDecode(t *testing.T) {
	dat, err := ioutil.ReadFile("./payload/merge_request-open.json")
	assert.Empty(t, err)

	mergeRequest := pkg.JSONDecode(string(dat))

	assert.Equal(t, "alexandre.eher", mergeRequest.User.Username)
}
