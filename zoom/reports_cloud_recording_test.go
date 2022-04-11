package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCloudRecording(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "",
			"to": "",
			"cloud_recording_storage": [
			  {
				"date": "",
				"free_usage": "",
				"plan_usage": "",
				"usage": ""
			  }
			]
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	cloud_recording, err := client.GetCloudRecording()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, cloud_recording, " response was empty")
	// assert.Equal(t, cloud_recording.CloudRecordingStorage[0].FreeUsage, "", "FreeUsage is not matching")
	// assert.Equal(t, cloud_recording.CloudRecordingStorage[0].PlanUsage, "", "PageSize is not matching")

}
