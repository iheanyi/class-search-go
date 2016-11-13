package main

import (
	_ "encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTermsService_ListTerms(t *testing.T) {
	termsBlob := `
	[
	{
		"code": "201620",
		"description": "Spring Semester 2017"
	},
	{
		"code": "201615",
		"description": "EMBA Fall 2016 (View Only)"
	}
	]
	`

	client, teardown := setup(func(mux *http.ServeMux) {
		mux.HandleFunc("/StudentRegistrationSsb/ssb/classSearch/getTerms",
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				fmt.Fprintf(w, termsBlob)
			})
	})

	defer teardown()

	terms, _, err := client.Terms.List()
	if err != nil {
		t.Errorf("Terms.List returned error: %v", err)
	}

	expected := []Term{
		{Code: "201620", Description: "Spring Semester 2017"},
		{Code: "201615", Description: "EMBA Fall 2016 (View Only)"},
	}

	if !reflect.DeepEqual(terms, expected) {
		t.Errorf("Terms.List returned %+v, expected %+v", terms, expected)
	}
}
