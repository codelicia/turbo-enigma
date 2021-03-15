package pkg

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

func PostJson(url string, message []byte) error {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(message))
	Assert(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := Client.Do(req)
	Assert(err)

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Response codeStatus code: %d, expected 200\n", resp.StatusCode))
	}

	return nil
}
