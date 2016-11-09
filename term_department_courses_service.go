package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	baseSearchPath = "StudentRegistrationSsb/ssb/searchResults/searchResults"
)

type TermDepartmentCoursesService struct {
	client *Client
}

func (ts *TermDepartmentCoursesService) List(term, dept string) ([]Course, *http.Response, error) {
	path := baseSearchPath

	req, err := ts.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Set("txt_subject", dept)
	q.Set("txt_term", term)
	q.Set("pageMaxSize", "500")
	req.URL.RawQuery = q.Encode()

	_, err = ts.client.AuthenticateClient(term)
	if err != nil {
		return nil, nil, err
	}

	var response Response
	res, err := ts.client.Do(req, &response)
	if err != nil {
		return nil, nil, err
	}

	// Decode Response.Data into a courses array
	var courses []Course
	coursesJson, err := json.Marshal(response.Data)
	if err != nil {
		return nil, nil, err
	}

	buf := bytes.NewBuffer(coursesJson)
	err = json.NewDecoder(buf).Decode(&courses)

	return courses, res, nil
}
