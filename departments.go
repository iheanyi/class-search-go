package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	departmentsURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/get_subject?searchTerm=&term=201620&offset=1&max=200"
)

func fetchDepartments() ([]Department, error) {
	fmt.Println("Starting to fetch departments")

	client, err := setupClient()

	if err != nil {
		return nil, err
	}

	r, err := doGet(client, departmentsURL)
	departments := make([]Department, 0)
	err = json.NewDecoder(strings.NewReader(r)).Decode(&departments)

	if err != nil {
		fmt.Println("Someting went wrong!")
		return nil, err
	}

	fmt.Println("Done fetching departments!")

	return departments, err
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
