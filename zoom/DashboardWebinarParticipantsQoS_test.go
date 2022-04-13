package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDashboardWebinarParticipants(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "",
			"page_count": "",
			"page_size": "",
			"total_records": "",
			"participants": [
			  {
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
					  "jitter": "12.9",
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
					  "frame_rate": "30 fps",
					  "resolution": ""
					}
				  } 
				],
				"version": "10.253"
			  }
			]
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	webinar_Participants, err := client.GetDashboardWebinarParticipants(" ")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, webinar_Participants, " response was empty")
	assert.NotNil(t, webinar_Participants.TotalRecords, "No records presents")
	assert.Equal(t, len(webinar_Participants.Participants), 1, "Total no of partcipants not matching")
	assert.Equal(t, len(webinar_Participants.Participants[0].Webinar_UserQos), 1, "Total no of partcipants not matching")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceFromCrc.AvgLoss, "20", "Average Loss of Webinar_AsDeviceFromCrc not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceToCrc.Bitrate, "1000", "Bitrate of Webinar_AsDeviceToCrc not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AsInput.Jitter, "12.5", "Jitter of Webinar_AsInput is not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.Latency, "20", "Latency of Webinar_AsOutput not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceFromCrc.MaxLoss, "10", "MaxLoss if Webinar_AudioDeviceFromCrc not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceToCrc.AvgLoss, "12", "Average Loss of Webinar_AudioDeviceToCrc not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AudioInput.Bitrate, "15", "Bitrate of Webinar_AudioInput not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_AudioOutput.Jitter, "12.9", "Jitter of Webinar_AudioOutput not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_CPUUsage.ZoomAvgCPUUsage, "60", "Avg CPU Usage of Webinar_CPUUsage not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceFromCrc.MaxLoss, "60", "Maxloss of Webinar_VideoDeviceFromCrc not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceToCrc.Latency, "16", "Latency of Webinar_VideoDeviceToCrc not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.Resolution, "720", "Resolution of Webinar_VideoInput not matched")
	assert.Equal(t, webinar_Participants.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.FrameRate, "30 fps", "Frame rate of Webinar_VideoOutput not matched")

}
