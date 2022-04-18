package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type SharingRecordingDetails struct {
	NextPageToken string       `json:"next_page_token"`
	PageCount     int          `json:"page_count"`
	PageSize      int          `json:"page_size"`
	TotalRecords  int          `json:"total_records"`
	Participants  Participants `json:"participants"`
}
type Participants []Participant
type Participant struct {
	ID       string  `json:"id"` //Metric
	UserID   string  `json:"user_id"`
	UserName string  `json:"user_name"`
	Details  Details `json:"details"`
}
type Details []Detail
type Detail struct {
	Content   string `json:"content"`    //Metric
	EndTime   string `json:"end_time"`   //Metric
	StartTime string `json:"start_time"` //Metric
}

func (z *Zoom) GetDashboardSharingRecordingDetails(webinar_id string) (SharingRecordingDetails, error) {

	url := fmt.Sprintf("/metrics/webinars/%v/participants/sharing", webinar_id)
	details, err := z.ReqBody("GET", url)
	if err != nil {
		log.Println("Error in GetSharingRecordingDetails response", err)
		return SharingRecordingDetails{}, err
	}

	var srdetail SharingRecordingDetails
	err = json.Unmarshal(details, &srdetail)
	if err != nil {
		log.Println("Error in GetSharingRecordingDetails, Unmarshalling JSON", err)
		return SharingRecordingDetails{}, err
	}

	return srdetail, err
}
