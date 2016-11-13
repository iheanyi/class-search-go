package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCourses_GetDescription(t *testing.T) {
	descriptionBlob := "This is some course description!"

	client, teardown := setup(func(mux *http.ServeMux) {
		mux.HandleFunc("/StudentRegistrationSsb/ssb/searchResults/getCourseDescription", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "POST")
			fmt.Fprintf(w, descriptionBlob)
		})
	})

	defer teardown()

	description, _, err := client.Courses.GetDescription("201620", "11131")

	if err != nil {
		t.Errorf("Courses.GetDescription returned error: %v", err)
	}

	expected := "This is some course description!"

	if description != expected {
		t.Errorf("Courses.GetDescription returned %+v, expected %+v", description, expected)
	}
}
