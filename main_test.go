package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func setup(config func(mux *http.ServeMux)) (*Client, func()) {
	mux := http.NewServeMux()
	config(mux)
	server := httptest.NewServer(mux)

	client, _ := NewClient(nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url

	return client, func() {
		server.Close()
	}
}

func testClientDefaultBaseURL(t *testing.T, c *Client) {
	if c.BaseURL == nil || c.BaseURL.String() != defaultBaseURL {
		t.Errorf("NewClient BaseURL = %v, expected %v", c.BaseURL, defaultBaseURL)
	}
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, execpted %v", r.Method, expected)
	}
}
