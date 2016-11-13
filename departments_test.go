package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestDepartmentsService_ListDepartments(t *testing.T) {
	departmentsBlob := `
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
		mux.HandleFunc("/StudentRegistrationSsb/ssb/classSearch/get_department",
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				fmt.Fprintf(w, departmentsBlob)
			})
	})

	defer teardown()

	term := &Term{
		Code:        "201620",
		Description: "Spring Semester 2017",
	}

	departments, _, err := client.Departments.List(term)
	if err != nil {
		t.Errorf("Departments.List returned error: %v", err)
	}

	expected := []Department{
		{Code: "ACCT", Name: "Accountancy"},
		{Code: "AME", Name: "Aerospace & Mechanical Engineering"},
	}

	if !reflect.DeepEqual(departments, expected) {
		t.Errorf("Departments.List returned %+v, expected %+v", departments, expected)
	}
}
