package main

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTermDepartmentCoursesService_ListTermDepartmentCourses(t *testing.T) {
	coursesBlob := `
	{
		"success": true,
		"totalCount": 65,
		"data": 
		[
		{
			"id": 295732,
			"term": "201620",
			"termDesc": "Spring Semester 2017",
			"courseReferenceNumber": "21976",
			"partOfTerm": "1",
			"courseNumber": "20100",
			"subject": "ACCT",
			"subjectDescription": "Accountancy",
			"sequenceNumber": "01",
			"campusDescription": "Main",
			"scheduleTypeDescription": "Class",
			"courseTitle": "Accountancy I",
			"creditHours": null,
			"maximumEnrollment": 11,
			"enrollment": 4,
			"seatsAvailable": 7,
			"waitCapacity": 0,
			"waitCount": 0,
			"waitAvailable": 0,
			"crossList": "66",
			"crossListCapacity": 45,
			"crossListCount": 4,
			"crossListAvailable": 41,
			"creditHourHigh": null,
			"creditHourLow": 3,
			"creditHourIndicator": null,
			"openSection": true,
			"linkIdentifier": null,
			"isSectionLinked": false,
			"subjectCourse": "ACCT20100",
			"faculty": 
			[
			{
				"bannerId": "901610394",
				"category": null,
				"class": "net.hedtech.banner.student.faculty.FacultyResultDecorator",
				"courseReferenceNumber": "21976",
				"displayName": "Larocque, Stephannie",
				"emailAddress": "larocque.1@nd.edu",
				"primaryIndicator": true,
				"term": "201620"
			}
			],
			"meetingsFaculty": 
			[
			{
				"category": "01",
				"class": "net.hedtech.banner.student.schedule.SectionSessionDecorator",
				"courseReferenceNumber": "21976",
				"faculty": 
				[
				],
				"meetingTime": {
					"beginTime": "1100",
					"building": "1144",
					"buildingDescription": "DeBartolo Hall",
					"campus": "M",
					"campusDescription": "Main",
					"category": "01",
					"class": "net.hedtech.banner.general.overall.MeetingTimeDecorator",
					"courseReferenceNumber": "21976",
					"creditHourSession": 3.0,
					"endDate": "05/03/2017",
					"endTime": "1215",
					"friday": false,
					"hoursWeek": 2.5,
					"meetingScheduleType": "CL",
					"monday": false,
					"room": "216",
					"saturday": false,
					"startDate": "01/17/2017",
					"sunday": false,
					"term": "201620",
					"thursday": true,
					"tuesday": true,
					"wednesday": false
				},
				"term": "201620"
			}
			]
		}
		],
		"pageOffset": 0,
		"pageMaxSize": 1,
		"sectionsFetchedCount": 65,
		"pathMode": "search",
		"searchResultsConfigs": null
	}
	`
	client, teardown := setup(func(mux *http.ServeMux) {
		mux.HandleFunc("/StudentRegistrationSsb/ssb/searchResults/searchResults",
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "GET")
				fmt.Fprintf(w, coursesBlob)
			})
	})

	defer teardown()

	var term string = "201620"
	var dept string = "ACCT"
	courses, _, err := client.TermDepartmentCourses.List(term, dept)
	if err != nil {
		t.Errorf("TermDepartmentCourses.List returned error: %v", err)
	}

	instructors := []CourseInstructor{
		{
			Name:      "Larocque, Stephannie",
			BannerId:  "901610394",
			Category:  "",
			IsPrimary: true,
		},
	}

	meeting := Meeting{
		StartTime:         "1100",
		EndTime:           "1215",
		BuildingId:        "1144",
		BuildingName:      "DeBartolo Hall",
		CreditHourSession: 3,
		CampusCode:        "M",
		RoomNumber:        "216",
		StartDate:         "01/17/2017",
		EndDate:           "05/03/2017",
		Sunday:            false,
		Monday:            false,
		Tuesday:           true,
		Wednesday:         false,
		Thursday:          true,
		Friday:            false,
		Saturday:          false,
	}

	meetings_faculty := []CourseMeetingFaculty{
		{
			Meeting: meeting,
		},
	}

	expected := []Course{
		{
			Id:                    295732,
			Term:                  "201620",
			Subject:               "ACCT",
			CourseReferenceNumber: "21976",
			SectionNumber:         "01",
			Title:                 "Accountancy I",
			CampusName:            "Main",
			IsOpen:                true,
			CrossList:             "66",
			CrossListCapacity:     45,
			CrossListAvailable:    41,
			CreditHourHigh:        0,
			CreditHourLow:         3,
			Instructors:           instructors,
			MeetingsFaculty:       meetings_faculty,
			SubjectCourse:         "ACCT20100",
		},
	}

	if !reflect.DeepEqual(courses, expected) {
		t.Errorf("TermDepartmentCourses.List returned %+v, expected %+v", courses, expected)
	}
}
