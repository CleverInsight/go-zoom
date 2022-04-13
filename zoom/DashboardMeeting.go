/* Author : Anusree TM
Created on : 8th April 2022
Description : This is a Module to Get meeting API from Dashboard Module ,
Documentation link - https://marketplace.zoom.us/docs/api-reference/zoom-api/dashboards/dashboardmeetings
*/
package zoom

import (
	"encoding/json"
	"log"
)

type DashboardMeeting struct {
	From          string   `json:"from"`
	To            string   `json:"to"`
	NextPageToken string   `json:"next_page_token"`
	PageCount     int      `json:"page_count"`
	PageSize      int      `json:"page_size"`
	TotalRecords  int      `json:"total_records"`
	Meetings      Meetings `json:"meetings"`
}
type Meetings []Meeting

type Meeting struct {
	Host               string         `json:"host"` // Metric
	Dept               string         `json:"dept"`
	Duration           string         `json:"duration"` // Metric of the form mm:ss
	Email              string         `json:"email"`
	EndTime            string         `json:"end_time"`
	Has3RdPartyAudio   bool           `json:"has_3rd_party_audio"` //Metric
	HasArchiving       bool           `json:"has_archiving"`       //Metric
	HasPstn            bool           `json:"has_pstn"`            //Metric
	HasRecording       bool           `json:"has_recording"`       //Metric
	HasScreenShare     bool           `json:"has_screen_share"`    //Metric
	HasSip             bool           `json:"has_sip"`             //Metric
	HasVideo           bool           `json:"has_video"`           //Metric
	HasVoip            bool           `json:"has_voip"`            //Metric
	ID                 int64          `json:"id"`                  // Metric
	InRoomParticipants string         `json:"in_room_participants"`
	Participants       int            `json:"participants"` //Metric
	StartTime          string         `json:"start_time"`
	Topic              string         `json:"topic"`
	UserType           string         `json:"user_type"`
	UUID               string         `json:"uuid"`
	AudioQuality       string         `json:"audio_quality"`        //Metric
	VideoQuality       string         `json:"video_quality"`        // Metric
	ScreenShareQuality string         `json:"screen_share_quality"` // Metric
	Key                keys           `json:"custom_keys"`
	TrackingFields     TrackingFields `json:"tracking_fields"`
}

type keys []Key
type TrackingFields []TrackingField
type Key struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TrackingField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func (z *Zoom) ListDashboardMeetings() (DashboardMeeting, error) {
	response, err := z.ReqBody("GET", "/metrics/meetings")
	if err != nil {
		log.Printf("Error during GetMeetings %v", err)
		return DashboardMeeting{}, err
	}

	// Unmarshal the response which is in byte format
	var meeting DashboardMeeting
	err = json.Unmarshal(response, &meeting)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return DashboardMeeting{}, err
	}

	return meeting, nil

}