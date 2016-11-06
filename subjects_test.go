package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSubjects_ListSubjects(t *testing.T) {
	subjectsBlob := `
	[
	{
		"code": "ACCT",
		"description": "Accountancy"
	},
	{
		"code": "AME",
		"description": "Aerospace & Mechanical Engineering"
	}
	]
	`

	client, teardown := setup(func(mux *http.ServeMux) {
		mux.HandleFunc("/StudentRegistrationSsb/ssb/classSearch/get_subject",
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				fmt.Fprintf(w, subjectsBlob)
			})
	})

	defer teardown()

	term := &Term{
		Code:        "201620",
		Description: "Spring Semester 2017",
	}

	subjects, _, err := client.Subjects.List(term)
	if err != nil {
		t.Errorf("Subjects.List returned error: %v", err)
	}

	expected := []Subject{
		{Code: "ACCT", Name: "Accountancy"},
		{Code: "AME", Name: "Aerospace & Mechanical Engineering"},
	}

	if !reflect.DeepEqual(subjects, expected) {
		t.Errorf("Subjects.List returned %+v, expected %+v", subjects, expected)
	}
}
