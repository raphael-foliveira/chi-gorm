package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
)

type testClient struct {
	ts *httptest.Server
}

func newTestClient(app *httptest.Server) *testClient {
	return &testClient{app}
}

func (tc *testClient) makeRequest(method string, endpoint string, body interface{}) (*http.Response, error) {
	hc := &http.Client{}
	url := tc.ts.URL + endpoint
	isNotGetOrDelete := !slices.Contains([]string{http.MethodGet, http.MethodDelete}, method)
	if body != nil && isNotGetOrDelete {
		return tc.sendRequestWithBody(hc, method, body, url)
	}
	return tc.sendRequest(hc, method, url)
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
