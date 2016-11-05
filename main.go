package main

import (
	"log"
)

const (
	sampleURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=ACCT&txt_term=201620&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc"
)

func main() {
	log.Print("Starting Program!")

	client, err := setupClient()
	if err != nil {
		log.Fatal(err)
	}

	authenticateClient(client, "201620")

	r, err := doGet(client, sampleURL)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(r)

	terms, err := FetchTerms()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(terms)

	_, err = FetchTermCourses(&terms[0])

	if err != nil {
		log.Fatal(err)
	}

	c, err := NewClient(nil)
	if err != nil {
		log.Fatal(err)
	}
	ts := &TermsService{client: c}

	terms, _, err = ts.List()
	if err != nil {
		log.Print("Something went wrong fetching terms.")
		log.Fatal(err)
	}
	log.Print(terms)

	log.Print("Done")
}
