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

func TestGetPlanUsage(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"plan_base": {
			  "hosts": "",
			  "type": "",
			  "usage": "",
			  "pending": 1
			},
			"plan_large_meeting": [
			  {
				"hosts": "",
				"type": "",
				"usage": "",
				"pending": 1
			  }
			],
			"plan_recording": {
			  "free_storage": "",
			  "free_storage_usage": "",
			  "plan_storage": "",
			  "plan_storage_exceed": "",
			  "plan_storage_usage": "",
			  "type": ""
			},
			"plan_united": {
			  "hosts": "",
			  "name": "",
			  "type": "",
			  "usage": "",
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
			  "hosts": "",
			  "type": "",
			  "usage": ""
			}
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	planusage, err := client.GetPlanUsage(" ")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, planusage, " response was empty")
	assert.Contains(t, planusage.PlanRecording.Type, "", "Type not matching")
	assert.Equal(t, planusage.PlanZoomRooms.Hosts, "", "Hosts not been matched")

}
