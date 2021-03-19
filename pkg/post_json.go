package pkg

import (
	"bytes"
	"fmt"
	"net/http"
)

func PostJSON(url string, message []byte) error {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(message))
	Assert(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := Client.Do(req)
	Assert(err)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("response codeStatus code: %d, expected 200", resp.StatusCode)
	}

	return nil
}
