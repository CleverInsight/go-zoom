package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportMeetings(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "w7587w4eiyfsudgk",
			"page_count": 1,
			"page_number": 1,
			"page_size": 30,
			"total_records": 20,
			"from": "2020-07-14",
			"meetings": [
			  {
				"custom_keys": [
				  {
					"key": "",
					"value": ""
				  }
				],
				"duration": 6,
				"end_time": "2020-07-15T23:30:19Z",
				"id": 12345,
				"participants_count": 2,
				"session_key": "ABC36jaBI145",
				"source": "Zoom",
				"start_time": "2019-07-15T23:24:52Z",
				"topic": "My Meeting",
				"total_minutes": 11,
				"type": 2,
				"user_email": "jchill@example.com",
				"user_name": "Jill Chill",
				"uuid": "4444AAAiAAAAAiAiAiiAii==",
				"schedule_time": "12/22/2021 16:20",
				"join_waiting_room_time": "02/11/2022 16:15",
				"join_time": "12/22/2021 16:20",
				"leave_time": "12/22/2021 17:13"
			  }
			],
			"to": "2020-08-14"
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	listmeetings, err := client.GetReportsMeetings(" ")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, listmeetings, " response was empty")
	assert.LessOrEqual(t, listmeetings.PageSize, 300, "Dept not matching")

}
