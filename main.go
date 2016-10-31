package main

import (
	"fmt"
	"log"
)

const (
	sampleURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=ACCT&txt_term=201620&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc"
)

func main() {
	fmt.Println("Starting Program!")

	client, err := setupClient()
	if err != nil {
		log.Fatal(err)
	}

	authenticateClient(client, "201620")

	r, err := doGet(client, sampleURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)

	terms, err := FetchTerms()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(terms)

	_, err = FetchTermCourses(&terms[0])

	if err != nil {
		log.Fatal(err)
	}

	professors, err := FetchInstructors(&terms[0])

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(professors)

	fmt.Println("Done")
}
