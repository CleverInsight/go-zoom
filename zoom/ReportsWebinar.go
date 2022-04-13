/* Author : Anusree TM
Created on : 8th April 2022
Description : This is a Module to Get the Webinar APi from reports module,
Documentation link - https://marketplace.zoom.us/docs/api-reference/zoom-api/methods/#operation/reportWebinarDetails
*/

package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Webinar struct {
	Dept             string        `json:"dept"`
	Duration         int           `json:"duration"` //Metric
	EndTime          string        `json:"end_time"`
	ID               string        `json:"id"`                 //Metric
	ParticipantCount int           `json:"participants_count"` //Metric
	StartTime        string        `json:"start_time"`
	Topic            string        `json:"topic"`         //Metric
	TotalTime        int           `json:"total_minutes"` //Metric
	Type             int           `json:"type"`          //Metric
	Email            string        `json:"user_email"`
	Name             string        `json:"user_name"`
	UUID             string        `json:"uuid"`
	WebinarKeys      WebinarKeys   `json:"webinar_keys"`
	WebinarFields    WebinarFields `json:"webinar_fields"`
}
type WebinarKeys []WebinarKey
type WebinarKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type WebinarFields []WebinarField
type WebinarField struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

//Return WebinarDetails Of Zoom
func (z *Zoom) GetWebinarDetails(webinar_id string) (Webinar, error) {
	url := fmt.Sprintf("/report/webinars/%v", webinar_id)
	response, err := z.ReqBody("GET", url)
	if err != nil {
		log.Printf("Error during GetWebinarDetails %v", err)
		return Webinar{}, err
	}

	// Unmarshal the response which is in byte format
	var webinar Webinar
	err = json.Unmarshal(response, &webinar)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return Webinar{}, err
	}

	return webinar, nil

}
