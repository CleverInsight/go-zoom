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
	PlanBase         PlanBase
	PlanLargeMeeting PlanLargeMeetings
	PlanRecording    PlanRecording
	PlanUnited       PlanUnited
	PlanWebinar      PlanWebinars
	PlanZoomEvents   PlanZoomEvents
	PlanZoomRooms    PlanZoomRooms
}
type PlanBase struct {
	Hosts   int    `json:"hosts"`
	Type    string `json:"type"`
	Usage   int    `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanLargeMeetings []PlanLargeMeeting
type PlanLargeMeeting struct {
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

type PlanWebinars []PlanWebinars
type PlanWebinar struct {
	Hosts   int    `json:"hosts"`
	Type    string `json:"type"`
	Usage   int    `json:"usage"`
	Pending int    `json:"pending"`
}
type PlanZoomEvents []PlanZoomEvent
type PlanZoomEvent struct {
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
