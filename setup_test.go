package fourth_time_attendance_test

import (
	"log"
	"net/url"
	"os"
	"testing"

	fourth_time_attendance "github.com/omniboost/go-fourth-time-attendance"
)

var (
	client *fourth_time_attendance.Client
)

func TestMain(m *testing.M) {
	var baseURL *url.URL
	var err error

	baseURLString := os.Getenv("FOURTH_BASE_URL")
	username := os.Getenv("FOURTH_USERNAME")
	password := os.Getenv("FOURTH_PASSWORD")

	client = fourth_time_attendance.NewClient(nil)
	client.SetUsername(username)
	client.SetPassword(password)

	if baseURLString != "" {
		baseURL, err = url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
	}

	if baseURL != nil {
		client.SetBaseURL(*baseURL)
	}

	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	m.Run()
}
