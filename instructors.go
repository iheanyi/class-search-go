package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	instructorURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/get_instructor?searchTerm=&offset=1&max=10000&term="
)

func FetchInstructors(t *Term) ([]Instructor, error) {
	fmt.Println("Starting to fetch instructors!")
	client, err := setupClient()

	if err != nil {
		return nil, err
	}

	fullSearchURL := instructorURL + t.Code

	authenticateClient(client, t.Code)
	fmt.Println(fullSearchURL)

	r, err := doGet(client, fullSearchURL)

	if err != nil {
		return nil, err
	}

	instructors := make([]Instructor, 0)
	err = json.NewDecoder(strings.NewReader(r)).Decode(&instructors)

	if err != nil {
		return nil, err
	}

	fmt.Println("Done fetching instructors!")
	return instructors, err
}
