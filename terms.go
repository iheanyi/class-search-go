package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Struct for the Course Return
type Course struct {
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

func FetchAllTermCourses() (string, error) {
	terms, err := FetchTerms()

	if err != nil {
		return "", err
	}

	for i := range terms {
		fmt.Println(terms[i].Code)
	}

	return "", err
}

func FetchTermDepartments(t *Term) ([]Department, error) {
	fullTermDepartmentURL := termDepartmentURL + t.Code

	client, err := setupClient()

	departments := make([]Department, 0)
	r, err := doGet(client, fullTermDepartmentURL)
	err = json.NewDecoder(strings.NewReader(r)).Decode(&departments)

	if err != nil {
		return nil, err
	}

	fmt.Println("Done Fetching departments for term: ", t.Code)

	return departments, err
}

// Fetch the courses for a subject in a specifc term.
func FetchTermDepartmentCourses(t *Term, d *Department) (string, error) {
	client, err := setupClient()
	if err != nil {
		return "", err
	}

	fullSearchURL := termDepartmentPart + d.Code + termStart + t.Code + searchURLEnd

	authenticateClient(client, t.Code)
	fmt.Println(fullSearchURL)
	r, err := doGet(client, fullSearchURL)

	if err != nil {
		return "", err
	}

	response := Response{}

	fmt.Println(response)

	// Marshal the response into it's relevant JSON.
	err = json.NewDecoder(strings.NewReader(r)).Decode(&response)
	if err != nil {
		return "", err
	}

	if !response.Success {
		log.Fatal("This request was unsuccessful!")
	}

	fmt.Println(response)
	fmt.Println(response.Data)

	return "", err
}

// Fetch all the courses for a term through the departments.
func FetchTermCourses(t *Term) (string, error) {
	client, err := setupClient()

	if err != nil {
		return "", err
	}

	// Authenticate client that we're using.
	authenticateClient(client, t.Code)

	// We need to Fetch all of the departments first.
	departments, err := FetchTermDepartments(t)
	fmt.Println(departments[0].Code)

	FetchTermDepartmentCourses(t, &departments[0])
	if err != nil {
		return "", err
	}

	return "", err
}
