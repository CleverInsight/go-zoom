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

func TestGetWebinarDetails(t *testing.T) {
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
			"duration": "",
			"end_time": "",
			"id": "",
			"participants_count": "",
			"start_time": "",
			"topic": "",
			"total_minutes": "",
			"tracking_fields": [
			  {
				"field": "",
				"value": ""
			  }
			],
			"type": "",
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
	assert.Equal(t, webinar.ID, "", "ID has not been matched")

}
