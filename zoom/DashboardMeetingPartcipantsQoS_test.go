package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDashboardMeetingParticipants(t *testing.T) {
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
				  "avg_loss": "20",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"as_device_to_crc": {
				  "avg_loss": "",
				  "bitrate": "1000",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"as_input": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "12.5",
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
				  "max_loss": "10"
				},
				"audio_device_to_crc": {
				  "avg_loss": "12",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"audio_input": {
				  "avg_loss": "",
				  "bitrate": "15",
				  "jitter": "",
				  "latency": "",
				  "max_loss": ""
				},
				"audio_output": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "12.5",
				  "latency": "",
				  "max_loss": ""
				},
				"cpu_usage": {
				  "system_max_cpu_usage": "",
				  "zoom_avg_cpu_usage": "60",
				  "zoom_max_cpu_usage": "",
				  "zoom_min_cpu_usage": ""
				},
				"date_time": "",
				"video_device_from_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": "60"
				},
				"video_device_to_crc": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "16",
				  "max_loss": ""
				},
				"video_input": {
				  "avg_loss": "",
				  "bitrate": "",
				  "jitter": "",
				  "latency": "",
				  "max_loss": "",
				  "frame_rate": "",
				  "resolution": "720"
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

	meeting_participants, err := client.GetDashboardMeetingParticipants(" ", "")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meeting_participants, " response was empty")
	assert.Equal(t, meeting_participants.UserID, "", "ID not matching")
	assert.Equal(t, len(meeting_participants.UserQos), 1, "Total no of partcipants not matching")
	assert.Equal(t, meeting_participants.UserQos[0].AsDeviceFromCrc.AvgLoss, "20", "Average Loss of AsDeviceFromCrc not matched")
	assert.Equal(t, meeting_participants.UserQos[0].AsDeviceToCrc.Bitrate, "1000", "Bitrate of AsDeviceToCrc not matched")
	assert.Equal(t, meeting_participants.UserQos[0].AsInput.Jitter, "12.5", "Jitter of AsInput is not matched")
	assert.Equal(t, meeting_participants.UserQos[0].AsOutput.Latency, "20", "Latency of AsOutput not matched")
	assert.Equal(t, meeting_participants.UserQos[0].AudioDeviceFromCrc.MaxLoss, "10", "MaxLoss if AudioDeviceFromCrc not matched")
	assert.Equal(t, meeting_participants.UserQos[0].AudioDeviceToCrc.AvgLoss, "12", "Average Loss of AudioDeviceToCrc not matched")
	assert.Equal(t, meeting_participants.UserQos[0].AudioInput.Bitrate, "15", "Bitrate of AudioInput not matched")
	assert.Equal(t, meeting_participants.UserQos[0].AudioOutput.Jitter, "12.5", "Jitter of AudioOutput not matched")
	assert.Equal(t, meeting_participants.UserQos[0].CPUUsage.ZoomAvgCPUUsage, "60", "Avg CPU Usage of CPUUsage not matched")
	assert.Equal(t, meeting_participants.UserQos[0].VideoDeviceFromCrc.MaxLoss, "60", "Maxloss of VideoDeviceFromCrc not matched")
	assert.Equal(t, meeting_participants.UserQos[0].VideoDeviceToCrc.Latency, "16", "Latency of VideoDeviceToCrc not matched")
	assert.Equal(t, meeting_participants.UserQos[0].VideoInput.Resolution, "720", "Resolution of VideoInput not matched")
	assert.Equal(t, meeting_participants.UserQos[0].VideoOutput.FrameRate, "20 fps", "Frame rate of VideoOutput not matched")

}
