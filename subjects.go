package main

import (
	"net/http"
)

const (
	subjectsBasePath = "StudentRegistrationSsb/ssb/classSearch/get_subject"
)

type Subject struct {
	Code string `json:"code"`
	Name string `json:"description"`
}

type SubjectsService struct {
	client *Client
}

func (ss *SubjectsService) List(t *Term) ([]Subject, *http.Response, error) {
	path := subjectsBasePath

	req, err := ss.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Set("term", t.Code)
	q.Set("offset", "1")
	q.Set("max", "200")
	req.URL.RawQuery = q.Encode()

	var subjects []Subject
	res, err := ss.client.Do(req, &subjects)

	if err != nil {
		return nil, nil, err
	}

	return subjects, res, nil
}
