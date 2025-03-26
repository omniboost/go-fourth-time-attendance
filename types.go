package fourth_time_attendance

import "encoding/xml"

var (
	ClockStatusCheckIn    ClockStatus = 1
	ClockStatusCheckOut   ClockStatus = 2
	ClockStatusBreakStart ClockStatus = 3
	ClockStatusBreakEnd   ClockStatus = 4
)

type ClockStatus int
type ClockingPostRequest struct {
	Root Root `xml:"Root"`
}

type Root struct {
	XMLName   xml.Name `xml:"Root"`
	Text      string   `xml:",chardata"`
	GroupGUID string   `xml:"GroupGUID,attr"`
	DateTime  DateTime `xml:"DateTime,attr"`
	Record    []Record `xml:"Record"`
}

type Record struct {
	XMLName       xml.Name    `xml:"Record"`
	Text          string      `xml:",chardata"`
	EmpNo         string      `xml:"EmpNo"`
	Location      string      `xml:"Location"`
	ClockStatus   ClockStatus `xml:"ClockStatus"`
	CheckIn       DateTime    `xml:"CheckIn"`
	CheckOut      DateTime    `xml:"CheckOut"`
	ActualMinutes string      `xml:"ActualMinutes"`
	Notes         string      `xml:"Notes"`
}

type ClockingPostResponse struct {
	XMLName          xml.Name `xml:"Result"`
	Text             string   `xml:",chardata"`
	Xsd              string   `xml:"xsd,attr"`
	Xsi              string   `xml:"xsi,attr"`
	DateTime         DateTime `xml:"DateTime,attr"`
	OrganisationId   string   `xml:"OrganisationId"`
	ProcessedRecords string   `xml:"ProcessedRecords"`
	SubmittedRecords string   `xml:"SubmittedRecords"`
	DeletedRecords   string   `xml:"DeletedRecords"`
}
