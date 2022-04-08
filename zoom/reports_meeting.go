package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type MeetingReports struct {
	NextPageToken string `json:"next_page_token"`
	PageCount     int    `json:"page_count"`
	PageNumber    int    `json:"page_number"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	From          string `json:"from"`
	Meetings      Mtings `json:"meetings"`
	To            string `json:"to"`
}
type Mtings []struct {
	Customkeys          CstmKeys  `json:"custom_keys"`
	Duration            int       `json:"duration"`
	EndTime             time.Time `json:"end_time"`
	ID                  int       `json:"id"`
	ParticipantsCount   int       `json:"participants_count"`
	SessionKey          string    `json:"session_key"`
	Source              string    `json:"source"`
	StartTime           time.Time `json:"start_time"`
	Topic               string    `json:"topic"`
	TotalMinutes        int       `json:"total_minutes"`
	Type                int       `json:"type"`
	UserEmail           string    `json:"user_email"`
	UserName            string    `json:"user_name"`
	UUID                string    `json:"uuid"`
	ScheduleTime        string    `json:"schedule_time"`
	JoinWaitingRoomTime string    `json:"join_waiting_room_time"`
	JoinTime            string    `json:"join_time"`
	LeaveTime           string    `json:"leave_time"`
}

type CstmKeys []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (z *Zoom) GetReportsMeetings(userId string) (MeetingReports, error) {
	url := fmt.Sprintf("report/users/%v/meetings", userId)
	report, err := z.ReqBody("GET", url)
	if err != nil {
		log.Println("Error in GetReportsMeetings response", err)
		return MeetingReports{}, err
	}

	var meetingreport MeetingReports

	err = json.Unmarshal(report, &meetingreport)
	if err != nil {
		log.Println("Error in GetReportsMeetings, Unmarshalling JSON ", err)
		return MeetingReports{}, err
	}

	return meetingreport, err
}
