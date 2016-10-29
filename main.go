package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

const (
	baseTermSearchURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/term/search?mode=search&term=201620"
	classSearchURL    = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/classSearch"
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

	r, err := doGet(client, baseTermSearchURL)
	if err != nil {
		log.Fatal(err)
	}

	r, err = doGet(client, sampleURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(r)
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
