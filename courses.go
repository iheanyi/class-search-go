package main

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
