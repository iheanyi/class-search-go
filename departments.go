package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	departmentsURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/get_subject?searchTerm=&term=201620&offset=1&max=200"
)

type Department struct {
	Code string `json:"code"`
	Name string `json:"description"`
}

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
