package main

import (
	"bytes"
	"net/http"
)

type Course struct {
	Id                       int                    `json:"id"`
	Term                     string                 `json:"term"`
	CourseRegistrationNumber string                 `json:"courseReferenceNumber"`
	Subject                  string                 `json:"subject"`
	SectionNumber            string                 `json:"sequenceNumber"`
	Title                    string                 `json:"courseTitle"`
	IsOpen                   bool                   `json:"openSection"`
	CrossList                string                 `json:"crossList"`
	CrossListCapacity        int                    `json:"crossListCapacity"`
	CrossListAvailable       int                    `json:"crossListAvailable"`
	CreditHourHigh           float64                `json:"creditHourHigh"`
	CreditHourLow            float64                `json:"creditHourLow"`
	Instructors              []CourseInstructor     `json:"faculty"`
	CampusName               string                 `json:"campusDescription"`
	MeetingsFaculty          []CourseMeetingFaculty `json:"meetingsFaculty"`
	SubjectCourse            string                 `json:"subjectCourse"`
}

type CourseInstructor struct {
	BannerId  string `json:"bannerId"`
	Category  string `json:"category"`
	IsPrimary bool   `json:"primaryIndicator"`
	Name      string `json:"displayName"`
}

type CourseMeetingFaculty struct {
	Meeting Meeting `json:"meetingTime"`
}

type Meeting struct {
	BuildingId        string  `json:"building"`
	BuildingName      string  `json:"buildingDescription"`
	CampusCode        string  `json:"campus"`
	CreditHourSession float64 `json:"creditHourSession"`
	EndDate           string  `json:"endDate"`
	EndTime           string  `json:"endTime"`
	RoomNumber        string  `json:"room"`
	StartDate         string  `json:"startDate"`
	StartTime         string  `json:"beginTime"`
	Sunday            bool    `json:"sunday"`
	Monday            bool    `json:"monday"`
	Tuesday           bool    `json:"tuesday"`
	Wednesday         bool    `json:"wednesday"`
	Thursday          bool    `json:"thursday"`
	Friday            bool    `json:"friday"`
	Saturday          bool    `json:"saturday"`
}

const (
	courseDescriptionPath = "StudentRegistrationSsb/ssb/searchResults/getCourseDescription"
)

type CoursesService struct {
	client *Client
}

func (cs *CoursesService) GetDescription(term, crn string) (string, *http.Response, error) {
	path := courseDescriptionPath

	req, err := cs.client.NewRequest("POST", path, nil)
	if err != nil {
		return "", nil, err
	}

	// It's URL form encoded, so we'll have that in here.
	q := req.URL.Query()
	q.Set("term", term)
	q.Set("courseReferenceNumber", crn)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")

	description := bytes.NewBuffer(nil)
	res, err := cs.client.Do(req, description)
	if err != nil {
		return "", res, err
	}

	return description.String(), res, nil
}
