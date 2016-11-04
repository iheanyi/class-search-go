package main

import (
	"encoding/json"
	"log"
	"strings"
)

type Course struct {
}

func FetchTermDepartmentCourses(t *Term, d *Department) (string, error) {
	client, err := setupClient()
	if err != nil {
		return "", err
	}

	fullSearchURL := termDepartmentPart + d.Code + termStart + t.Code + searchURLEnd

	authenticateClient(client, t.Code)
	log.Print(fullSearchURL)
	r, err := doGet(client, fullSearchURL)

	if err != nil {
		return "", err
	}

	response := Response{}

	log.Print(response)

	// Marshal the response into it's relevant JSON.
	err = json.NewDecoder(strings.NewReader(r)).Decode(&response)
	if err != nil {
		return "", err
	}

	if !response.Success {
		log.Fatal("This request was unsuccessful!")
	}

	log.Print(response)
	log.Print(response.Data)

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
	log.Print(departments[0].Code)

	FetchTermDepartmentCourses(t, &departments[0])
	if err != nil {
		return "", err
	}

	return "", err
}

func FetchAllTermCourses() (string, error) {
	terms, err := FetchTerms()

	if err != nil {
		return "", err
	}

	for _, term := range terms {
		log.Print(term.Code)
	}

	return "", err
}
