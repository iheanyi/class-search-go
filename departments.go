package main

import (
	"net/http"
)

type Department struct {
	Code string `json:"code"`
	Name string `json:"description"`
}

const (
	departmentsBasePath = "StudentRegistrationSsb/ssb/classSearch/get_department"
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
	req.URL.RawQuery = q.Encode()

	var departments []Department
	res, err := ds.client.Do(req, &departments)
	if err != nil {
		return nil, nil, err
	}

	return departments, res, nil
}
