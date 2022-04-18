/* Author : Reethu Reymond
Created on : 8th April 2022
Description : This is a Unit test  module the Billing_plan_usage.go
*/

package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBillingPlanUsage(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"plan_base": {
			  "hosts": "20",
			  "type": "Business",
			  "usage": "20",
			  "pending": 1
			},
			"plan_large_meeting": [
			  {
				"hosts": "20",
				"type": "Business",
				"usage": "20",
				"pending": 1
			  }
			],
			"plan_recording": {
			  "free_storage": "",
			  "free_storage_usage": "",
			  "plan_storage": "",
			  "plan_storage_exceed": "",
			  "plan_storage_usage": "",
			  "type": "Business"
			},
			"plan_united": {
			  "hosts": "",
			  "name": "",
			  "type": "",
			  "usage": "20",
			  "pending": 1
			},
			"plan_webinar": [
			  {
				"hosts": "",
				"type": "",
				"usage": "",
				"pending": 1
			  }
			],
			"plan_zoom_events": [
			  {
				"hosts": "",
				"type": "",
				"usage": "",
				"pending": 1
			  }
			],
			"plan_zoom_rooms": {
			  "hosts": "20",
			  "type": "",
			  "usage": ""
			}
		  }`))
	}))
	client := &Zoom{
		Endpoint: testServer.URL,
	}
	planusage, err := client.GetBillingPlanUsage(" ")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, planusage, " response was empty")
	assert.Equal(t, planusage.PlanBase.Hosts, "20", "Hosts has not been matched")
	assert.Equal(t, planusage.PlanBase.Pending, 1, "Pending has not been matched")
	assert.Contains(t, planusage.PlanBase.Type, "Business", "Type has not been matched")
	assert.Contains(t, planusage.PlanRecording.Type, "Business", "Type not matching")
	assert.Equal(t, planusage.PlanUnited.Usage, "20", "Type has not been matched")
	assert.Equal(t, planusage.PlanZoomRooms.Hosts, "20", "Hosts has not been matched")
	assert.Equal(t, planusage.PlanLargeMeeting[0].Hosts, "20", "Hosts has not been matched")
	assert.Equal(t, planusage.PlanZoomEvents[0].Pending, 1, "Pending has not been matched")
}
