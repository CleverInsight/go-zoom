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

type PlanUsage struct {
	PlanBase         PlanBase
	PlanLargeMeeting PlanLargeMeeting
	PlanRecording    PlanRecording
	PlanUnited       PlanUnited
	PlanWebinar      PlanWebinar
	PlanZoomEvents   PlanZoomEvents
	PlanZoomRooms    PlanZoomRooms
}
type PlanBase struct {
	Hosts   string `json:"hosts"`
	Type    string `json:"type"`
	Usage   string `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanLargeMeeting []struct {
	Hosts   string `json:"hosts"`
	Type    string `json:"type"`
	Usage   string `json:"usage"`
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
	Hosts   string `json:"hosts"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Usage   string `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanWebinar []struct {
	Hosts   string `json:"hosts"`
	Type    string `json:"type"`
	Usage   string `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanZoomEvents []struct {
	Hosts   string `json:"hosts"`
	Type    string `json:"type"`
	Usage   string `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanZoomRooms struct {
	Hosts string `json:"hosts"`
	Type  string `json:"type"`
	Usage string `json:"usage"`
}

//Return PlanUsage Of Zoom
func (z *Zoom) GetPlanUsage(account_id string) (PlanUsage, error) {

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
