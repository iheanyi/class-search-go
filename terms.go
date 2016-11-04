package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	termSearchURL      = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/getTerms?searchTerm=&offset=1&max=10"
	termDepartmentPart = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject="
	termStart          = "&txt_term="
	searchURLEnd       = "&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc"
	termDepartmentURL  = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/get_subject?searchTerm=&offset=1&max=500&term="
)

// Struct for each term
type Term struct {
	Code        string `json:"code,number"`
	Description string `json:"description"`
}

type Response struct {
	Success    bool                     `json:"success"`
	TotalCount int                      `json:"totalCount"`
	Data       []map[string]interface{} `json:"data"`
}

// FetchTerms fetches an API response of all of the most recent terms for the
// Notre Dame Class Search API.
func FetchTerms() ([]Term, error) {
	fmt.Println("Starting to Fetch terms!")

	client, err := setupClient()

	if err != nil {
		return nil, err
	}

	r, err := doGet(client, termSearchURL)

	terms := make([]Term, 0)
	//err = json.Unmarshal([]byte(r), &terms)
	err = json.NewDecoder(strings.NewReader(r)).Decode(&terms)

	if err != nil {
		fmt.Println("Something went wrong!")
		return nil, err
	}

	fmt.Println("Done Fetching terms!")
	return terms, err
}
