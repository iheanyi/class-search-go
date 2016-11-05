package main

import (
	"log"
	"net/http"
)

const (
	departmentsBasePath = "StudentRegistrationSsb/ssb/classSearch/get_subject"
)

type DepartmentsService struct {
	client *Client
}

func (ds *DepartmentsService) List(t *Term) ([]Department, *http.Response, error) {
	path := departmentsBasePath

	req, err := ds.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Set("term", t.Code)
	q.Set("offset", "1")
	q.Set("max", "200")
	var departments []Department
	res, err := ds.client.Do(req, &departments)

	log.Print(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return departments, res, nil
}
