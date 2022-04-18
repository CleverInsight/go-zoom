package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportsCloudRecording(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "",
			"to": "",
			"cloud_recording_storage": [
			  {
				"date": "",
				"free_usage": "1024MB",
				"plan_usage": "1024MB",
				"usage": "1.62GB"
			  }
			]
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	cloud_recording, err := client.GetReportsCloudRecording()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, cloud_recording, " response was empty")
	assert.Equal(t, cloud_recording.CloudRecordingStorage[0].FreeUsage, "1024MB", "FreeUsage is not matching")
	assert.Equal(t, cloud_recording.CloudRecordingStorage[0].PlanUsage, "1024MB", "PlanUsage is not matching")
	assert.Equal(t, cloud_recording.CloudRecordingStorage[0].Usage, "1.62GB", "Usage is not matching")

}
