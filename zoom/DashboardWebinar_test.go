package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListDashboardWebinar(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "",
			"to": "",
			"next_page_token": "Tva2CuIdTgsv8wAnhyAdU3m06Y2HuLQtlh3",
			"page_count": 1,
			"page_size": 300,
			"total_records": 1,
			"webinars": [
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
				"has_3rd_party_audio": false,
				"has_archiving": true,
				"has_pstn": false,
				"has_recording": false,
				"has_screen_share": true,
				"has_sip": false,
				"has_video": false,
				"has_voip": true,
				"id": 668464,
				"participants": 20,
				"start_time": "",
				"topic": "Zoom Test Cases",
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

	listwebinar, err := client.ListDashboardWebinar()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, listwebinar, " response was empty")
	assert.Contains(t, listwebinar.DashboardWebinar[0].Dept, "", "Dept not matching")
	assert.Equal(t, listwebinar.PageCount, 1, "PageCount not been matched")
	assert.LessOrEqual(t, listwebinar.PageSize, 300, "Pagesize not been matched")
	assert.Equal(t, listwebinar.DashboardWebinar[0].Has3RdPartyAudio, false, "does not have 3rd party audio")
	assert.Equal(t, listwebinar.DashboardWebinar[0].Participants, 20, "Participant Count not matching")
	assert.Equal(t, listwebinar.DashboardWebinar[0].Topic, "Zoom Test Cases", "Topic not matching")

}
