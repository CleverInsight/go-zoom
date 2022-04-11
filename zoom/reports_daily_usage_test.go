package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportsDailyUsage(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"dates": [
			  {
				"date": "",
				"meeting_minutes": "",
				"meetings": "",
				"new_users": "",
				"participants": ""
			  }
			],
			"month": "",
			"year": ""
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	dailyusage, err := client.GetReportsDailyUsage()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, dailyusage, " response was empty")
	//assert.Equal(t, dailyusage.Dates[0].MeetingMinutes, 1, "FreeUsage is not matching")
	//assert.Equal(t, cloud_recording.CloudRecordingStorage[0].PlanUsage, "", "PageSize is not matching")

}
