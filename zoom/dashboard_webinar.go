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
	Duration           string `json:"duration"` //Metric
	Email              string `json:"email"`
	EndTime            string `json:"end_time"`
	Has3RdPartyAudio   bool   `json:"has_3rd_party_audio"` //Metric
	HasArchiving       bool   `json:"has_archiving"`       //Metric
	HasPstn            bool   `json:"has_pstn"`            //Metric
	HasRecording       bool   `json:"has_recording"`       //Metric
	HasScreenShare     bool   `json:"has_screen_share"`    //Metric
	HasSip             bool   `json:"has_sip"`             //Metric
	HasVideo           bool   `json:"has_video"`           //Metric
	HasVoip            bool   `json:"has_voip"`            //Metric
	ID                 int64  `json:"id"`                  //Metric
	Participants       int    `json:"participants"`        //Metric
	StartTime          string `json:"start_time"`
	Topic              string `json:"topic"` //Metric
	UserType           string `json:"user_type"`
	UUID               string `json:"uuid"`
	AudioQuality       string `json:"audio_quality"`
	VideoQuality       string `json:"video_quality"`
	ScreenShareQuality string `json:"screen_share_quality"`
	Webinar            Webinar
}
type Webinars []struct {
	Host       string `json:"host"` //Metric
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
