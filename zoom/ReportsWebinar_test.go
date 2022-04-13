/* Author : Reethu Reymond
Created on : 8th April 2022
Description : This is a Unit test  module the reports_webinar.go
*/
package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportsWebinarDetails(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"custom_keys": [
			  {
				"key": "",
				"value": ""
			  }
			],
			"dept": "",
			"duration": 20,
			"end_time": "",
			"id": "2456112",
			"participants_count": 12,
			"start_time": "",
			"topic": "abc",
			"total_minutes": 30,
			"tracking_fields": [
			  {
				"field": "",
				"value": ""
			  }
			],
			"type": 20,
			"user_email": "",
			"user_name": "",
			"uuid": ""
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	webinar, err := client.GetWebinarDetails(" ")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, webinar, " results was empty")
	//assert.NotNil(t, webinar.Keys, "", "Keys are not matching")
	assert.Equal(t, webinar.Duration, 20, "Duration has not been matched")
	assert.Equal(t, webinar.ID, "2456112", "ID has not been matched")
	assert.Equal(t, webinar.Topic, "abc", "Topic has not been matched")
	assert.Equal(t, webinar.TotalTime, 30, "TotalTime has not been matched")
	assert.Equal(t, webinar.Type, 20, "Type has not been matched")

}
