package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func makeRequest(testServerUrl string) func(method string, endpoint string, body any) (*http.Response, error) {
	return func(method string, endpoint string, body any) (*http.Response, error) {
		hc := &http.Client{}
		url := testServerUrl + endpoint
		if body != nil {
			return sendRequestWithBody(hc, method, body, url)
		}
		return sendRequest(hc, method, url)
	}
}

func sendRequest(hc *http.Client, method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return hc.Do(req)
}

func sendRequestWithBody(hc *http.Client, method string, body any, url string) (*http.Response, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, err
	}
	return hc.Do(req)
}
