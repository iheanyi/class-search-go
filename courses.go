package main

type Course struct {
	Id                       int     `json:"id"`
	Term                     string  `json:"term"`
	CourseRegistrationNumber string  `json:"courseReferenceNumber"`
	Subject                  string  `json:"subject"`
	SectionNumber            string  `json:"sequenceNumber"`
	Title                    string  `json:"courseTitle"`
	IsOpen                   bool    `json:"openSection"`
	CrossList                string  `json:"crossList"`
	CrossListCapacity        int     `json:"crossListCapacity"`
	CrossListAvailable       int     `json:"crossListAvailable"`
	CreditHourHigh           float64 `json:"creditHourHigh"`
	CreditHourLow            float64 `json:"creditHourLow"`
}
