package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	baseURL           = "https://ssb.cc.nd.edu"
	baseTermSearchURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/term/termSelection?mode=search"
	classSearchURL    = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/classSearch"
	termSearchURL     = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/classSearch/getTerms?searchTerm=&offset=1&max=10"
	sampleURL         = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=ACCT&txt_term=201620&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc"
)

func authenticateClient(client *http.Client) []*http.Cookie {
	resp, err := client.Get(baseTermSearchURL)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, respBodyErr := ioutil.ReadAll(resp.Body)

	if respBodyErr != nil {
		log.Fatal(err)
	}

	return resp.Cookies()
}

func sendTermRequest(client *http.Client) *http.Response {
	searchResp, err := client.Get(termSearchURL)
	defer searchResp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	_, searchErr := ioutil.ReadAll(searchResp.Body)

	if searchErr != nil {
		log.Fatal(searchErr)
	}

	return searchResp
}

func main() {
	fmt.Println("Starting Program!")

	cookieJar, _ := cookiejar.New(nil)
	u, err := url.Parse(baseURL)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: cookieJar,
	}

	cookies := authenticateClient(client)

	u, err = url.Parse(sampleURL)

	if err != nil {
		log.Fatal(err)
	}

	client.Jar.SetCookies(u, cookies)

	resp, err := client.Get(sampleURL)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Shouldn't be empty :o
	fmt.Println(string(body))
	fmt.Println(cookieJar.Cookies(u))
	fmt.Println(client.Jar)
}
