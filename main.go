package main

import (
	"log"
)

const (
	sampleURL = "https://ssb.cc.nd.edu/StudentRegistrationSsb/ssb/searchResults/searchResults?txt_subject=ACCT&txt_term=201620&startDatepicker=&endDatepicker=&pageOffset=0&pageMaxSize=10&sortColumn=subjectDescription&sortDirection=asc"
)

func main() {
	log.Print("Starting Program!")

	c, err := NewClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	terms, _, err := c.Terms.List()
	if err != nil {
		log.Print("Something went wrong fetching terms.")
		log.Fatal(err)
	}
	log.Print(terms)

	departments, _, err := c.Departments.List(&terms[0])
	if err != nil {
		log.Print("Something went wrong fetching departments.")
		log.Fatal(err)
	}
	log.Print(departments)

	subjects, _, err := c.Subjects.List(&terms[0])
	if err != nil {
		log.Print("Something went wrong fetching subjects!")
		log.Fatal(err)
	}
	log.Print(subjects)

	courses, _, err := c.TermDepartmentCourses.List(terms[0].Code, departments[0].Code)
	if err != nil {
		log.Print("Something went wrong fetching courses.")
		log.Fatal(err)
	}
	log.Print(courses)

	courses, _, err = c.TermDepartmentCourses.List(terms[0].Code, departments[1].Code)
	if err != nil {
		log.Print("Something went wrong fetching courses.")
		log.Fatal(err)
	}
	log.Print(courses)

	instructors, _, err := c.Instructors.List(terms[0].Code)
	if err != nil {
		log.Print("Something went wrong fetching instructors.")
		log.Fatal(err)
	}
	log.Print(instructors)

	//_, err = FetchTermCourses(&terms[0])
	log.Print("Done")
}
