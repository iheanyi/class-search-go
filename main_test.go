package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func setup(config func(mux *http.ServeMux)) (*Client, func()) {
	mux := http.NewServeMux()
	mux.HandleFunc("/"+baseAuthPath, func(w http.ResponseWriter, r *http.Request) {
		cookie := &http.Cookie{Name: "JSESSIONID", Value: "foobarbaz"}
		http.SetCookie(w, cookie)
	})

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

var tests = []struct {
	name            string
	mock            func(mux *http.ServeMux)
	check           func(t *testing.T, client *Client)
	skipIntegration bool
}{
	{
		name:            "authenticate",
		mock:            func(mux *http.ServeMux) {},
		check:           testAuthenticateClient,
		skipIntegration: true,
	},
}

func TestEverythingMock(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, teardown := setup(tt.mock)
			defer teardown()

			tt.check(t, client)
		})
	}
}

func TestEverythingIntegration(t *testing.T) {
	if os.Getenv("TEST_INTEGRATION") == "" {
		t.Skipf("not running test against real API")
	}

	for _, tt := range tests {
		if tt.skipIntegration {
			continue
		}
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(nil)
			if err != nil {
				t.Fatal(err)
			}
			tt.check(t, client)
		})
	}
}

func TestAuthenticateClient(t *testing.T) {
	client, teardown := setup(func(_ *http.ServeMux) {})
	defer teardown()
	testAuthenticateClient(t, client)
}

func TestAuthenticateClientIntegration(t *testing.T) {
	if os.Getenv("TEST_INTEGRATION") == "" {
		t.Skipf("not running test against real API")
	}

	client, _ := NewClient(nil)
	testAuthenticateClient(t, client)
}

func testClientDefaults(t *testing.T, c *Client) {
	testClientDefaultBaseURL(t, c)
}

func TestNewClient(t *testing.T) {
	c, _ := NewClient(nil)
	testClientDefaults(t, c)
}

func testAuthenticateClient(t *testing.T, client *Client) {
	res, err := client.AuthenticateClient("201620")
	if err != nil {
		t.Fatalf("TestAuthenicate returned an error: %v", err)
	}

	if want, got := 1, len(res.Cookies()); want != got {
		t.Fatalf("want len %d, got %d", want, got)
	}
	cookie := res.Cookies()[0]
	if want, got := "JSESSIONID", cookie.Name; want != got {
		t.Errorf("want=%q", want)
		t.Errorf(" got=%q", got)
	}
}
