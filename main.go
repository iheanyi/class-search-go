package main

import (
	"log"
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

	courses, _, err := c.TermDepartmentCourses.List(terms[0].Code, subjects[0].Code)
	if err != nil {
		log.Print("Something went wrong fetching courses.")
		log.Fatal(err)
	}
	log.Print(courses)

	courses, _, err = c.TermDepartmentCourses.List(terms[0].Code, subjects[1].Code)
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
	log.Print(courses[0])
	description, _, err := c.Courses.GetDescription(terms[0].Code, courses[0].CourseReferenceNumber)
	if err != nil {
		log.Print("Something went wrong fetching course description.")
		log.Fatal(err)
	}
	log.Printf("%+v", description)
	log.Print(description)

	log.Print("Done")
}
