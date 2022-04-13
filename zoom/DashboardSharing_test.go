package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDashboardSharingRecordingDetails(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
      "next_page_token": "Tva2CuIdTgsv8wAnhyAdU3m06Y2HuLQtlh3",
      "page_count": 1,
      "page_size": 30,
      "total_records": 1,
      "participants": [
        {
          "details": [
            {
              "content": "bookmark",
              "end_time": "",
              "start_time": ""
            }
          ],
          "id": "USR111",
          "user_id": "",
          "user_name": ""
        }
      ]
    }`))
	}))
	client := &Zoom{
		Endpoint: testServer.URL,
	}
	dashboard_sharing, err := client.GetSharingRecordingDetails("")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, dashboard_sharing, " response was empty")
	assert.Equal(t, dashboard_sharing.Participants[0].ID, "USR111", "ID not been matched")
	assert.Equal(t, dashboard_sharing.Participants[0].Details[0].Content, "bookmark", "Content not been matched")
	assert.Equal(t, dashboard_sharing.Participants[0].Details[0].EndTime, "", "EndTime not been matched")
	assert.Equal(t, dashboard_sharing.Participants[0].Details[0].StartTime, "", "StartTime not been matched")
}
