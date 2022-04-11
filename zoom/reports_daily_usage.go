package zoom

import (
	"encoding/json"
	"log"
)

type DailyUsage struct {
	Dates Dates
	Month string `json:"month"`
	Year  string `json:"year"`
}
type Dates []Date
type Date struct {
	Date           string `json:"date"`
	MeetingMinutes string `json:"meeting_minutes"`
	Meetings       string `json:"meetings"`
	NewUsers       string `json:"new_users"`
	Participants   string `json:"participants"`
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
