package zoom

import (
	"encoding/json"
	"log"
)

type DashBoardWebinar struct {
	From               string `json:"from"`
	To                 string `json:"to"`
	NextPageToken      string `json:"next_page_token"`
	PageCount          int    `json:"page_count"`
	PageSize           int    `json:"page_size"`
	TotalRecords       int    `json:"total_records"`
	Dept               string `json:"dept"`
	Duration           string `json:"duration"`
	Email              string `json:"email"`
	EndTime            string `json:"end_time"`
	Has3RdPartyAudio   bool   `json:"has_3rd_party_audio"`
	HasArchiving       string `json:"has_archiving"`
	HasPstn            string `json:"has_pstn"`
	HasRecording       string `json:"has_recording"`
	HasScreenShare     string `json:"has_screen_share"`
	HasSip             string `json:"has_sip"`
	HasVideo           string `json:"has_video"`
	HasVoip            string `json:"has_voip"`
	ID                 string `json:"id"`
	Participants       string `json:"participants"`
	StartTime          string `json:"start_time"`
	Topic              string `json:"topic"`
	UserType           string `json:"user_type"`
	UUID               string `json:"uuid"`
	AudioQuality       string `json:"audio_quality"`
	VideoQuality       string `json:"video_quality"`
	ScreenShareQuality string `json:"screen_share_quality"`
	Webinar            Webinar
}
type Webinars []struct {
	Host       string `json:"host"`
	CustomKeys CustomKeys
}
type CustomKeys []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Return ListWebinar Of Zoom
func (z *Zoom) ListWebinar() (DashBoardWebinar, error) {

	response, err := z.ReqBody("GET", "/metrics/webinars")
	if err != nil {
		log.Printf("Error during ListWebinar %v", err)
		return DashBoardWebinar{}, err
	}

	// Unmarshal the response which is in byte format
	var webinar DashBoardWebinar
	err = json.Unmarshal(response, &webinar)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return DashBoardWebinar{}, err
	}

	return webinar, nil

}
