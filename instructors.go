package main

import (
	"net/http"
)

type Instructor struct {
	BannerId string `json:"code"`
	Name     string `json:"description"`
}

type InstructorsService struct {
	client *Client
}

const (
	instructorsBasePath = "StudentRegistrationSsb/ssb/classSearch/get_instructor"
)

func (is *InstructorsService) List(term string) ([]Instructor, *http.Response, error) {
	path := instructorsBasePath

	req, err := is.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Set("term", term)
	q.Set("offset", "1")
	q.Set("max", "5000")
	req.URL.RawQuery = q.Encode()
	var instructors []Instructor
	res, err := is.client.Do(req, &instructors)

	if err != nil {
		return nil, nil, err
	}

	return instructors, res, nil
}
