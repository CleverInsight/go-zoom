package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeetingParticipants(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"device": "",
			"domain": "",
			"harddisk_id": "",
			"ip_address": "",
			"join_time": "",
			"leave_time": "",
			"location": "",
			"mac_addr": "",
			"pc_name": "",
			"user_id": "",
			"user_name": "",
			"user_qos": [
			  {
				"as_device_from_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"as_device_to_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"as_input": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": "",
				  "frame_rate": "",
				  "resolution": ""
				},
				"as_output": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "20",
				  "max_loss": "",
				  "frame_rate": "",
				  "resolution": ""
				},
				"audio_device_from_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"audio_device_to_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"audio_input": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"audio_output": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"cpu_usage": {
				  "system_max_cpu_usage": "",
				  "zoom_avg_cpu_usage": "",
				  "zoom_max_cpu_usage": "",
				  "zoom_min_cpu_usage": ""
				},
				"date_time": "",
				"video_device_from_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"video_device_to_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"video_input": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": "",
				  "frame_rate": "",
				  "resolution": ""
				},
				"video_output": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": "",
				  "frame_rate": "20 fps",
				  "resolution": ""
				}
			  }
			],
			"version": ""
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	meeting_participants, err := client.GetMeetingParticipants(" ", "")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meeting_participants, " response was empty")
	assert.Equal(t, meeting_participants.UserID, "", "ID not matching")
}
