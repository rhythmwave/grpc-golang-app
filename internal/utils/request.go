package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CallAPI(method string, url string, body []byte) ([]byte, error) {
	// 1. Create a new HTTP client
	client := &http.Client{}

	// 2. Create a new HTTP request
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// 3. Set headers (optional)
	// You can set additional headers here if needed
	// req.Header.Set("Content-Type", "application/json") // Example

	// 4. Do the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Close the response body

	// 5. Handle the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API call failed with status code: %d", resp.StatusCode)
	}

	// 6. Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
