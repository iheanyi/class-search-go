package main

import (
	"log"
	"net/http"
)

const (
	termsBasePath = "StudentRegistrationSsb/ssb/classSearch/getTerms"
)

type TermsService struct {
	client *Client
}

func (ts *TermsService) List() ([]Term, *http.Response, error) {
	path := termsBasePath

	req, err := ts.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Set("searchTerm", "")
	q.Set("offset", "1")
	q.Set("max", "10")
	req.URL.RawQuery = q.Encode()

	var terms []Term
	res, err := ts.client.Do(req, &terms)

	log.Print(res.Body)
	if err != nil {
		log.Print("Something went wrong!")
		return nil, nil, err
	}
	return terms, res, nil
}
