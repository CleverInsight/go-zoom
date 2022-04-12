package zoom

import (
	"encoding/json"
	"log"
)

//new users, num of meetings, participant count and total meeting mins).
type DailyUsage struct {
	Dates Dates
	Month int `json:"month"`
	Year  int `json:"year"`
}
type Dates []Date
type Date struct {
	Date           string `json:"date"`
	MeetingMinutes int    `json:"meeting_minutes"` //Metric
	Meetings       int    `json:"meetings"`        //Metric
	NewUsers       int    `json:"new_users"`       //Metric
	Participants   int    `json:"participants"`    //Metric
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
