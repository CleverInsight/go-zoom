package zoom

import (
	"encoding/json"
	"log"
)

//new users, num of meetings, participant count and total meeting mins).
type DailyUsage struct {
	Dates Dates  `json:"dates"`
	Month string `json:"month"`
	Year  string `json:"year"`
}
type Dates []Date
type Date struct {
	Date           string `json:"date"`
	MeetingMinutes string `json:"meeting_minutes"` //Metric
	Meetings       string `json:"meetings"`        //Metric
	NewUsers       string `json:"new_users"`       //Metric
	Participants   string `json:"participants"`    //Metric
}

func (z *Zoom) GetReportsDailyUsage() (DailyUsage, error) {
	usage, err := z.ReqBody("GET", "/report/daily")
	if err != nil {
		log.Println("Error in GetReportsDailyUsage response", err)
		return DailyUsage{}, err
	}

	var dailyusage DailyUsage

	err = json.Unmarshal(usage, &dailyusage)
	if err != nil {
		log.Println("Error in GetReportsDailyUsage, Unmarshallling the json", err)
		return DailyUsage{}, err
	}

	return dailyusage, err
}
