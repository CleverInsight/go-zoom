/* Author : Anusree TM
Created on : 8th April 2022
Description : This is a Module to Get the Plan Usage from the API,
Documentation link - https://marketplace.zoom.us/docs/api-reference/zoom-api/ma/#operation/getPlanUsage
*/

package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

//All  the Structs in the below struct are Metrics
type PlanUsage struct {
	PlanBase         PlanBase         `json:"plan_base"`
	PlanLargeMeeting PlanLargeMeeting `json:"plan_large_meeting"`
	PlanRecording    PlanRecording    `json:"plan_recording"`
	PlanUnited       PlanUnited       `json:"plan_united"`
	PlanWebinar      PlanWebinar      `json:"plan_webinar"`
	PlanZoomEvents   PlanZoomEvents   `json:"plan_zoom_events"`
	PlanZoomRooms    PlanZoomRooms    `json:"plan_zoom_rooms"`
}
type PlanBase struct {
	Hosts   int    `json:"hosts"`
	Type    string `json:"type"`
	Usage   int    `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanLargeMeeting []BillingPlanLargeMeeting

type BillingPlanLargeMeeting struct {
	Hosts   int    `json:"hosts"`
	Type    string `json:"type"`
	Usage   int    `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanRecording struct {
	FreeStorage       string `json:"free_storage"`
	FreeStorageUsage  string `json:"free_storage_usage"`
	PlanStorage       string `json:"plan_storage"`
	PlanStorageExceed string `json:"plan_storage_exceed"`
	PlanStorageUsage  string `json:"plan_storage_usage"`
	Type              string `json:"type"`
}
type PlanUnited struct {
	Hosts   int    `json:"hosts"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Usage   int    `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanWebinar []BillingPlanWebinar

type BillingPlanWebinar struct {
	Hosts   int    `json:"hosts"`
	Type    string `json:"type"`
	Usage   int    `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanZoomEvents []BillingPlanZoomEvents

type BillingPlanZoomEvents struct {
	Hosts   int    `json:"hosts"`
	Type    string `json:"type"`
	Usage   int    `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanZoomRooms struct {
	Hosts int    `json:"hosts"`
	Type  string `json:"type"`
	Usage int    `json:"usage"`
}

//Return PlanUsage Of Zoom
func (z *Zoom) GetBillingPlanUsage(account_id string) (PlanUsage, error) {

	url := fmt.Sprintf("/accounts/%v/plans/usage", account_id)
	response, err := z.ReqBody("GET", url)
	if err != nil {
		log.Printf("Error during GetUsers %v", err)
		return PlanUsage{}, err
	}

	// Unmarshal the response which is in byte format
	var planusage PlanUsage
	err = json.Unmarshal(response, &planusage)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return PlanUsage{}, err
	}

	return planusage, nil

}
