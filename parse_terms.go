package main

import (
	_ "encoding/json"
	"fmt"
)

const (
	termSearchURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/getTerms?searchTerm=&offset=1&max=10"
)

type Term struct {
	Code        string
	Description string
}

// FetchTerms fetches an API response of all of the most recent terms for the
// Notre Dame Class Search API.
func fetchTerms(target interface{}) (string, error) {
	fmt.Println("Starting to fetch terms!")

	client, err := setupClient()

	if err != nil {
		return "", err
	}

	r, err := doGet(client, termSearchURL)
	if err != nil {
		return "", err
	}

	fmt.Println("Done fetching terms!")
	//return json.NewDecoder(r.Body).Decode(target), err
	return r, err
}
