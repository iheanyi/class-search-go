package main

import (
	"net/http"
)

const (
	baseSearchPath = "ssb/searchResults/searchResults"
)

type TermDepartmentCoursesService struct {
	client *Client
}

func (ts *TermDepartmentCoursesService) List(t *Term, d *Department) ([]Course, *http.Response, error) {
	path := baseSearchPath

	req, err := ts.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Set("txt_subject", d.Code)
	q.Set("txt_term", t.Code)
	q.Set("pageMazSize", "500")

	var courses []Course
	res, err := ts.client.Do(req, &courses)
	if err != nil {
		return nil, nil, err
	}

	return courses, res, nil
}
