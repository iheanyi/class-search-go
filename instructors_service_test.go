package main

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestInstructorsService_ListInstructors(t *testing.T) {
	instructorsBlob := `
	[
		{
			"code": "123456",
			"description": "Instructor One"
		},
		{
			"code": "543210",
			"description": "Instructor Two"
		}
	]
	`

	client, teardown := setup(func(mux *http.ServeMux) {
		mux.HandleFunc("/StudentRegistrationSsb/ssb/classSearch/get_instructor", func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, "GET")
			fmt.Fprintf(w, instructorsBlob)
		})
	})

	defer teardown()

	instructors, _, err := client.Instructors.List("201620")
	if err != nil {
		t.Errorf("Instructors.List returned error: %v", err)
	}

	expected := []Instructor{
		{BannerId: "123456", Name: "Instructor One"},
		{BannerId: "543210", Name: "Instructor Two"},
	}

	if !reflect.DeepEqual(instructors, expected) {
		t.Errorf("Instructors.List returned %+v, expected %+v", instructors, expected)
	}
}
