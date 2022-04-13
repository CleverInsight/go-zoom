package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListDashboardMeeting(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "",
			"to": "",
			"next_page_token": "Tva2CuIdTgsv8wAnhyAdU3m06Y2HuLQtlh3",
			"page_count": 2,
			"page_size": 30,
			"total_records": 1,
			"meetings": [
			  {
				"host": "",
				"custom_keys": [
				  {
					"key": "",
					"value": ""
				  }
				],
				"dept": "",
				"duration": "15:30",
				"email": "",
				"end_time": "",
				"has_3rd_party_audio": true,
				"has_archiving": true,
				"has_pstn": true,
				"has_recording": true,
				"has_screen_share": true,
				"has_sip": true,
				"has_video": true,
				"has_voip": true,
				"id": 221651,
				"in_room_participants": "",
				"participants": 20,
				"start_time": "",
				"topic": "",
				"tracking_fields": [
				  {
					"field": "",
					"value": ""
				  }
				],
				"user_type": "",
				"uuid": "",
				"audio_quality": "",
				"video_quality": "",
				"screen_share_quality": ""
			  }
			]
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	meetings, err := client.ListDashboardMeetings()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meetings, " response was empty")
	assert.Equal(t, meetings.PageCount, 2, "PageCount is not matching")
	assert.Equal(t, meetings.PageSize, 30, "PageSize is not matching")
	assert.Equal(t, meetings.Meetings[0].Has3RdPartyAudio, true, "Does not have 3rd party audio")
	assert.Equal(t, meetings.Meetings[0].Duration, "15:30", "Duration does not match")

}
