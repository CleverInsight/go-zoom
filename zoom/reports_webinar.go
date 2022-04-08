/* Author : Anusree TM
Created on : 4th April 2022
Description : This is a Module to Get the Plan Usage from the API,
Documentation link - https://marketplace.zoom.us/docs/api-reference/zoom-api/methods/#operation/reportWebinarDetails
*/

package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Webinar struct {
	Dept             string `json:"dept"`
	Duration         string `json:"duration"`
	EndTime          string `json:"end_time"`
	ID               string `json:"id"`
	ParticipantCount string `json:"participants_count"`
	StartTime        string `json:"start_time"`
	Topic            string `json:"topic"`
	TotalTime        string `json:"total_minutes"`
	Type             string `json:"type"`
	Email            string `json:"user_email"`
	Name             string `json:"user_name"`
	UUID             string `json:"uuid"`
	Keys             Keys
	Fields           Fields
}
type Keys []struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Fields []struct {
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
