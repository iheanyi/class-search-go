package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

const (
	termSearchURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/getTerms?searchTerm=&offset=1&max=10"
)

// Each term per semester
type Term struct {
	Code        string
	Description string
}

// FetchTerms fetches an API response of all of the most recent terms for the
// Notre Dame Class Search API.
func fetchTerms() ([]Term, error) {
	fmt.Println("Starting to fetch terms!")

	client, err := setupClient()

	if err != nil {
		return nil, err
	}

	r, err := doGet(client, termSearchURL)

	terms := make([]Term, 0)
	err = json.NewDecoder(strings.NewReader(r)).Decode(&terms)

	if err != nil {
		fmt.Println("Something went wrong!")
		log.Fatal(err)
		return nil, err
	}

	fmt.Println("Done fetching terms!")
	return terms, err
}
