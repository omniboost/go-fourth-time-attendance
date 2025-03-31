package fourth_time_attendance_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	fourth_time_attendance "github.com/omniboost/go-fourth-time-attendance"
)

func TestPostClockRequest(t *testing.T) {
	req := client.NewPostClockRequest()
	req.PathParams().ClientID = "lpdemothirty"
	req.RequestBody().Root.GroupGUID = "ACDC1805-90E9-C27F-5467-4763F3C19B19"
	req.RequestBody().Root.DateTime = fourth_time_attendance.DateTime{time.Date(2025, 03, 26, 8, 30, 0, 0, time.UTC)}
	req.RequestBody().Root.Record = []fourth_time_attendance.Record{
		{EmpNo: "9999999999", Location: "SCDEMO1", ClockStatus: fourth_time_attendance.ClockStatusCheckIn, CheckIn: fourth_time_attendance.DateTime{time.Date(2025, 03, 25, 8, 30, 0, 0, time.UTC)}, CheckOut: fourth_time_attendance.DateTime{}, ActualMinutes: "0", Notes: ""},
		{EmpNo: "255", Location: "SCDEMO1", ClockStatus: fourth_time_attendance.ClockStatusCheckOut, CheckIn: fourth_time_attendance.DateTime{}, CheckOut: fourth_time_attendance.DateTime{time.Date(2025, 03, 26, 18, 30, 0, 0, time.UTC)}, ActualMinutes: "0", Notes: ""},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
