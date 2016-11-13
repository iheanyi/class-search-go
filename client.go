package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	defaultBaseURL  = "https://ssb.cc.nd.edu/"
	baseAuthPath    = "StudentRegistrationSsb/ssb/term/search"
	baseAuthTermURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/term/search?mode=search&term="
)

type Response struct {
	Success    bool                     `json:"success"`
	TotalCount int                      `json:"totalCount"`
	Data       []map[string]interface{} `json:"data"`
}

type Client struct {
	client  *http.Client
	BaseURL *url.URL

	Terms                 *TermsService
	Departments           *DepartmentsService
	TermDepartmentCourses *TermDepartmentCoursesService
	Courses               *CoursesService
	Subjects              *SubjectsService
	Instructors           *InstructorsService
}

func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		cookieJar, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}

		httpClient = &http.Client{
			Jar: cookieJar,
		}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.Terms = &TermsService{client: c}
	c.TermDepartmentCourses = &TermDepartmentCoursesService{client: c}
	c.Departments = &DepartmentsService{client: c}
	c.Subjects = &SubjectsService{client: c}
	c.Instructors = &InstructorsService{client: c}
	return c, nil
}

// In order to authenticate a client, you have to authenticate it with a term.
// What this ends up doing is adding a JSESSIONID to the cookie jar.
func (c *Client) AuthenticateClient(term string) (*http.Response, error) {
	path := baseAuthPath

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("mode", "search")
	q.Set("term", term)
	req.URL.RawQuery = q.Encode()

	res, err := c.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if rerr := resp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			buf := bytes.NewBuffer(nil)
			io.Copy(buf, resp.Body)
			err := json.NewDecoder(buf).Decode(v)
			if err != nil {
				return nil, err
			}
		}
	}

	return resp, err
}

func setupClient() (*http.Client, error) {
	cookieJar, err := cookiejar.New(nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: cookieJar,
	}

	return client, err
}

func authenticateClient(c *http.Client, term string) {
	authURL := baseAuthTermURL + term
	log.Print(authURL)

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
