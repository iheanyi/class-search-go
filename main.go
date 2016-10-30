package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

const (
	baseTermSearchURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/term/search?mode=search&term="
	sampleURL         = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=ACCT&txt_term=201620&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc"
)

func main() {
	fmt.Println("Starting Program!")

	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: cookieJar,
	}

	authenticateClient(client, "201620")

	r, err := doGet(client, sampleURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)

	terms, err := fetchTerms()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(terms)
	fmt.Println("Done")
}

func setupClient() (*http.Client, error) {
	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	client := &http.Client{
		Jar: cookieJar,
	}

	return client, err
}

func authenticateClient(c *http.Client, term string) {
	s := []string{baseTermSearchURL, term}
	authURL := strings.Join(s, "")
	fmt.Println(authURL)

	_, err := doGet(c, authURL)
	if err != nil {
		log.Fatal(err)
	}
}

func doGet(c *http.Client, URL string) (string, error) {
	resp, err := c.Get(URL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(response), err
}
