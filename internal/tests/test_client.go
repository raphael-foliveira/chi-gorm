package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

type testClient struct {
	ts *httptest.Server
}

func (t *testClient) makeRequest(method string, endpoint string, body interface{}) (*http.Response, error) {
	hc := &http.Client{}
	url := t.ts.URL + endpoint
	if body != nil {
		return t.sendRequestWithBody(hc, method, body, url)
	}
	return t.sendRequest(hc, method, url)
}

func (t *testClient) sendRequest(hc *http.Client, method string, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	return hc.Do(req)
}

func (t *testClient) sendRequestWithBody(hc *http.Client, method string, body interface{}, url string) (*http.Response, error) {
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
