/* Author : Anusree TM
Created on : 8th April 2022
Description : This is a Module to Get the Meeting participants QoS API from Dashboard module  ,
Documentation link - https://marketplace.zoom.us/docs/api-reference/zoom-api/methods/#operation/dashboardMeetingParticipantQOS
*/
package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type MeetingParticipantsQos struct {
	Device     string `json:"device"`
	Domain     string `json:"domain"`
	HarddiskID string `json:"harddisk_id"`
	IP         string `json:"ip_address"`
	JoinTime   string `json:"join_time"`
	LeaveTime  string `json:"leave_time"`
	Location   string `json:"location"`
	MacAddress string `json:"mac_addr"`
	PcName     string `json:"pc_name"`
	UserID     string `json:"user_id"`
	Name       string `json:"user_name"`
	DateTime   string `json:"date_time"`
	Version    string `json:"version"`
	UserQos    UserQos
}
type UserQos []Userqos
type Userqos struct {
	AsDeviceFromCrc    AsDeviceFromCrc
	AsInput            AsInput
	AsOutput           AsOutput
	AudioDeviceFromCrc AudioDeviceFromCrc
	AudioDeviceToCrc   AudioDeviceToCrc
	AudioInput         AudioInput
	AudioOutput        AudioOutput
	CPUUsage           CPUUsage
	VideoDeviceFromCrc VideoDeviceFromCrc
	VideoDeviceToCrc   VideoDeviceToCrc
	VideoInput         VideoInput
	VideoOutput        VideoOutput
}

// All the below structs are Metrics
type AsDeviceFromCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type AsDeviceToCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type AsInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}
type AsOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}
type AudioDeviceFromCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type AudioDeviceToCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type AudioInput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type AudioOutput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type CPUUsage struct {
	SystemMaxCPUUsage string `json:"system_max_cpu_usage"`
	ZoomAvgCPUUsage   string `json:"zoom_avg_cpu_usage"`
	ZoomMaxCPUUsage   string `json:"zoom_max_cpu_usage"`
	ZoomMinCPUUsage   string `json:"zoom_min_cpu_usage"`
}
type VideoDeviceFromCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type VideoDeviceToCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type VideoInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}
type VideoOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}

//Return MeetingParticipants Of Zoom
func (z *Zoom) GetMeetingParticipants(meeting_id string, participant_id string) (MeetingParticipantsQos, error) {
	url := fmt.Sprintf("/metrics/meetings/%v/participants/%v/qos", meeting_id, participant_id)
	response, err := z.ReqBody("GET", url)
	if err != nil {
		log.Printf("Error during GetMeetingParticipants %v", err)
		return MeetingParticipantsQos{}, err
	}

	// Unmarshal the response which is in byte format
	var meeting_participants MeetingParticipantsQos
	err = json.Unmarshal(response, &meeting_participants)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return MeetingParticipantsQos{}, err
	}

	return meeting_participants, nil

}
