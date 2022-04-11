/* Author : Anusree TM
Created on : 4th April 2022
Description : This is a Module to Get the Webinar Participants QoS from Dashboard Module,
Documentation link - https://marketplace.zoom.us/docs/api-reference/zoom-api/methods/#operation/dashboardMeetingParticipantQOS
*/

package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type WebinarParticipants struct {
	NextPageToken string `json:"next_page_token"`
	PageCount     string `json:"page_count"`
	PageSize      string `json:"page_size"`
	TotalRecords  string `json:"total_records"`
	Participants  Participts
}
type Participts []struct {
	Device          string `json:"device"`
	Domain          string `json:"domain"`
	HarddiskID      string `json:"harddisk_id"`
	IPAddress       string `json:"ip_address"`
	JoinTime        string `json:"join_time"`
	LeaveTime       string `json:"leave_time"`
	Location        string `json:"location"`
	MacAddr         string `json:"mac_addr"`
	PcName          string `json:"pc_name"`
	UserID          string `json:"user_id"`
	UserName        string `json:"user_name"`
	Webinar_UserQos Webinar_UserQos
}
type Webinar_UserQos []struct {
	DateTime                   string `json:"date_time"`
	Version                    string `json:"version"`
	Webinar_AsDeviceFromCrc    Webinar_AsDeviceFromCrc
	Webinar_AsInput            Webinar_AsInput
	Webinar_AudioDeviceFromCrc Webinar_AudioDeviceFromCrc
	Webinar_AudioDeviceToCrc   Webinar_AudioDeviceToCrc
	Webinar_AudioInput         Webinar_AudioInput
	Webinar_AudioOutput        Webinar_AudioOutput
	Webinar_CPUUsage           Webinar_CPUUsage
	Webinar_VideoDeviceFromCrc Webinar_VideoDeviceFromCrc
	Webinar_VideoDeviceToCrc   Webinar_VideoDeviceToCrc
	Webinar_VideoInput         Webinar_VideoInput
	Webinar_VideoOutput        Webinar_VideoOutput
}
type Webinar_AsDeviceFromCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_AsDeviceToCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_AsInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}
type Webinar_AsOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}
type Webinar_AudioDeviceFromCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_AudioDeviceToCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_AudioInput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_AudioOutput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_CPUUsage struct {
	SystemMaxCPUUsage string `json:"system_max_cpu_usage"`
	ZoomAvgCPUUsage   string `json:"zoom_avg_cpu_usage"`
	ZoomMaxCPUUsage   string `json:"zoom_max_cpu_usage"`
	ZoomMinCPUUsage   string `json:"zoom_min_cpu_usage"`
}
type Webinar_VideoDeviceFromCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_VideoDeviceToCrc struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Webinar_VideoInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}
type Webinar_VideoOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	FrameRate  string `json:"frame_rate"`
	Resolution string `json:"resolution"`
}

//Return WebinarParticipants Of Zoom
func (z *Zoom) GetWebinarParticipants(webinar_id string) (WebinarParticipants, error) {
	url := fmt.Sprintf("/metrics/webinars/%v/participants/qos", webinar_id)
	response, err := z.ReqBody("GET", url)
	if err != nil {
		log.Printf("Error during GetWebinarParticipants %v", err)
		return WebinarParticipants{}, err
	}

	// Unmarshal the response which is in byte format
	var webinar_participants WebinarParticipants
	err = json.Unmarshal(response, &webinar_participants)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return WebinarParticipants{}, err
	}

	return webinar_participants, nil

}
