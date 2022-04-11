package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDashListWebinar(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "",
			"to": "",
			"next_page_token": "Tva2CuIdTgsv8wAnhyAdU3m06Y2HuLQtlh3",
			"page_count": 1,
			"page_size": 400,
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
				"has_archiving": "",
				"has_pstn": "",
				"has_recording": "",
				"has_screen_share": "",
				"has_sip": "",
				"has_video": "",
				"has_voip": "",
				"id": "",
				"participants": 20,
				"start_time": "",
				"topic": "",
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

	listwebinar, err := client.ListWebinar()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, listwebinar, " response was empty")
	assert.Contains(t, listwebinar.Dept, "", "Dept not matching")
	assert.Equal(t, listwebinar.PageCount, 1, "PageCount not been matched")
	assert.Equal(t, listwebinar.PageSize, 400, "Pagesize not been matched")
	assert.Equal(t, listwebinar.Has3RdPartyAudio, false, "does not have 3rd party audio")

}
