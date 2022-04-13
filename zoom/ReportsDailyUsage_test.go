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
				"meeting_minutes": "20",
				"meetings": "9",
				"new_users": "12",
				"participants": "20"
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
	assert.Equal(t, dailyusage.Dates[0].MeetingMinutes, "20", "MeetingMinute is not matching")
	assert.Equal(t, dailyusage.Dates[0].Meetings, "9", "Meetingsis not matching")
	assert.Equal(t, dailyusage.Dates[0].NewUsers, "12", "NewUser is not matching")
	assert.Equal(t, dailyusage.Dates[0].Participants, "20", "Participants is not matching")
}
