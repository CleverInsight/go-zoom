package zoom

import (
	"encoding/json"
	"log"
)

type DashBoardMeeting struct {
	From          string `json:"from"`
	To            string `json:"to"`
	NextPageToken string `json:"next_page_token"`
	PageCount     int    `json:"page_count"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	Meetings      Meetings
}
type Meetings []struct {
	Host               string `json:"host"`
	Dept               string `json:"dept"`
	Duration           string `json:"duration"`
	Email              string `json:"email"`
	EndTime            string `json:"end_time"`
	Has3RdPartyAudio   string `json:"has_3rd_party_audio"`
	HasArchiving       string `json:"has_archiving"`
	HasPstn            string `json:"has_pstn"`
	HasRecording       string `json:"has_recording"`
	HasScreenShare     string `json:"has_screen_share"`
	HasSip             string `json:"has_sip"`
	HasVideo           string `json:"has_video"`
	HasVoip            string `json:"has_voip"`
	ID                 string `json:"id"`
	InRoomParticipants string `json:"in_room_participants"`
	Participants       string `json:"participants"`
	StartTime          string `json:"start_time"`
	Topic              string `json:"topic"`
	UserType           string `json:"user_type"`
	UUID               string `json:"uuid"`
	AudioQuality       string `json:"audio_quality"`
	VideoQuality       string `json:"video_quality"`
	ScreenShareQuality string `json:"screen_share_quality"`
	Key                Key
	TrackingFields     TrackingFields
}
type Key []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type TrackingFields []struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

func (z *Zoom) GetMeetings() (DashBoardMeeting, error) {
	response, err := z.ReqBody("GET", "/metrics/meetings")
	if err != nil {
		log.Printf("Error during GetMeetings %v", err)
		return DashBoardMeeting{}, err
	}

	// Unmarshal the response which is in byte format
	var meeting DashBoardMeeting
	err = json.Unmarshal(response, &meeting)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return DashBoardMeeting{}, err
	}

	return meeting, nil

}
