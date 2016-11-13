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

func (ts *TermDepartmentCoursesService) List(term, subject string) ([]Course, *http.Response, error) {
	path := baseSearchPath

	req, err := ts.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	// With this request, every single request has to hit the re-authentication
	// endpoint because each department call is otherwise cached by the cookie. :/
	res, err := ts.client.AuthenticateClient(term)
	if err != nil {
		return nil, res, err
	}

	q := req.URL.Query()
	q.Set("txt_subject", subject)
	q.Set("txt_term", term)
	q.Set("pageMaxSize", "500")
	req.URL.RawQuery = q.Encode()

	var response Response
	res, err = ts.client.Do(req, &response)
	if err != nil {
		return nil, res, err
	}

	// Decode Response.Data into a courses array
	var courses []Course
	coursesJson, err := json.Marshal(response.Data)
	if err != nil {
		return nil, nil, err
	}

	buf := bytes.NewBuffer(coursesJson)
	err = json.NewDecoder(buf).Decode(&courses)
	if err != nil {
		return nil, nil, err
	}

	return courses, res, nil
}
