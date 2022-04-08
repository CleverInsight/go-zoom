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
	Participant   Participants `json:"participants"`
}
type Participants []struct {
	Detail   Details `json:"details"`
	ID       string  `json:"id"`
	UserID   string  `json:"user_id"`
	UserName string  `json:"user_name"`
}

type Details []struct {
	Content   string `json:"content"`
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

func (z *Zoom) GetSharingRecordingDetails(webinar_id string) (SharingRecordingDetails, error) {

	url := fmt.Sprintf("metrics/webinars/%v/participants/sharing", webinar_id)
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
