package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeeting(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "",
			"to": "",
			"next_page_token": "Tva2CuIdTgsv8wAnhyAdU3m06Y2HuLQtlh3",
			"page_count": 1,
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
				"duration": "",
				"email": "",
				"end_time": "",
				"has_3rd_party_audio": "",
				"has_archiving": "",
				"has_pstn": "",
				"has_recording": "",
				"has_screen_share": "",
				"has_sip": "",
				"has_video": "",
				"has_voip": "",
				"id": "",
				"in_room_participants": "",
				"participants": "",
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

	meetings, err := client.GetMeetings()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meetings, " response was empty")
	assert.Equal(t, meetings.PageCount,1, "PageCount is not matching")
	assert.Equal(t, meetings.PageSize,30, "PageSize is not matching")

}