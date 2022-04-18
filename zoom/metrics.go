package zoom

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Durationsplit(a string) float64 {
	Durationsplit := strings.Split(a, ":")

	return (StringtoFloat(Durationsplit[0]) * 60) + StringtoFloat(Durationsplit[1])
}
func Booltofloat(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func StringtoFloat(a string) float64 {
	floatans, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Println(err)
	}
	return floatans
}
func (z *Zoom) GetZoomMetrics() (map[string]float64, error) {

	headers := make(map[string][]string)
	token := "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6ImlyVVRpT01kUmVHQ3dfeDJiV3ZNSHciLCJleHAiOjE2ODA5NTIyMDAsImlhdCI6MTY0OTQxMDg0NX0.n12WeYogkzxCQTjl495Fpzx-k3k15f-Jz5EfyuTb5pY"
	headers["Authorization"] = []string{token}

	client := Zoom{
		ApiKey:   "irUTiOMdReGCw_x2bWvMHw",
		ApiToken: "r3JMzIsHn41NcxVEXmlE7EF8f4J9OmJxCarn",
		Headers:  headers,
		Endpoint: "https://api.zoom.us/v2/",
	}
	ZoomMetrics := make(map[string]float64)
	// Saranyan Wrote the below

	// Meeting Participants API Client
	Meetingparticipantsclient, err := client.GetDashboardMeetingParticipants("", "")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are metrics of Meeting Participants QoS from Dashboard API")

	// AsDeviceFromCrc
	ZoomMetrics["AsDeviceFromCrc - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceFromCrc.AvgLoss)
	ZoomMetrics["AsDeviceFromCrc - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceFromCrc.Bitrate)
	ZoomMetrics["AsDeviceFromCrc - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceFromCrc.Jitter)
	ZoomMetrics["AsDeviceFromCrc - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceFromCrc.Latency)
	ZoomMetrics["AsDeviceFromCrc - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceFromCrc.MaxLoss)
	//AsDeviceToCrc
	ZoomMetrics["AsDeviceToCrc - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceToCrc.AvgLoss)
	ZoomMetrics["AsDeviceToCrc - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceToCrc.Bitrate)
	ZoomMetrics["AsDeviceToCrc - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceToCrc.Jitter)
	ZoomMetrics["AsDeviceToCrc - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceToCrc.Latency)
	ZoomMetrics["AsDeviceToCrc - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsDeviceToCrc.MaxLoss)
	//AsInput
	ZoomMetrics["AsInput - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsInput.AvgLoss)
	ZoomMetrics["AsInput - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsInput.Bitrate)
	ZoomMetrics["AsInput - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsInput.Jitter)
	ZoomMetrics["AsInput - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsInput.Latency)
	ZoomMetrics["AsInput - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsInput.MaxLoss)
	ZoomMetrics["AsInput - Resolution"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsInput.Resolution)
	ZoomMetrics["AsInput - FrameRate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsInput.FrameRate)
	//AsOutput
	ZoomMetrics["AsOutput - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsOutput.AvgLoss)
	ZoomMetrics["AsOutput - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsOutput.Bitrate)
	ZoomMetrics["AsOutput - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsOutput.Jitter)
	ZoomMetrics["AsOutput - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsOutput.Latency)
	ZoomMetrics["AsOutput - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsOutput.MaxLoss)
	ZoomMetrics["AsOutput - Resolution"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsOutput.Resolution)
	ZoomMetrics["AsOutput - FrameRate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AsOutput.FrameRate)
	// AudioDeviceFromCrc
	ZoomMetrics["AudioDeviceFromCrc - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceFromCrc.AvgLoss)
	ZoomMetrics["AudioDeviceFromCrc - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceFromCrc.Bitrate)
	ZoomMetrics["AudioDeviceFromCrc - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceFromCrc.Jitter)
	ZoomMetrics["AudioDeviceFromCrc - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceFromCrc.Latency)
	ZoomMetrics["AudioDeviceFromCrc - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceFromCrc.MaxLoss)
	// AudioDeviceToCrc
	ZoomMetrics["AudioDeviceToCrc - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceToCrc.AvgLoss)
	ZoomMetrics["AudioDeviceToCrc - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceToCrc.Bitrate)
	ZoomMetrics["AudioDeviceToCrc - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceToCrc.Jitter)
	ZoomMetrics["AudioDeviceToCrc - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceToCrc.Latency)
	ZoomMetrics["AudioDeviceToCrc - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioDeviceToCrc.MaxLoss)
	// AudioInput
	ZoomMetrics["AudioInput - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioInput.AvgLoss)
	ZoomMetrics["AudioInput - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioInput.Bitrate)
	ZoomMetrics["AudioInput - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioInput.Jitter)
	ZoomMetrics["AudioInput - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioInput.Latency)
	ZoomMetrics["AudioInput - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioInput.MaxLoss)
	// AudioOutput
	ZoomMetrics["AudioOutput - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioOutput.AvgLoss)
	ZoomMetrics["AudioOutput - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioOutput.Bitrate)
	ZoomMetrics["AudioOutput - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioOutput.Jitter)
	ZoomMetrics["AudioOutput - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioOutput.Latency)
	ZoomMetrics["AudioOutput - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].AudioOutput.MaxLoss)
	//CPUUsage
	ZoomMetrics["CPUUsage - SystemMaxCPUUsage"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].CPUUsage.SystemMaxCPUUsage)
	ZoomMetrics["CPUUsage - ZoomAvgCPUUsage"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].CPUUsage.ZoomAvgCPUUsage)
	ZoomMetrics["CPUUsage - ZoomMaxCPUUsage"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].CPUUsage.ZoomMaxCPUUsage)
	ZoomMetrics["CPUUsage - ZoomMinCPUUsage"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].CPUUsage.ZoomMinCPUUsage)

	// VideoDeviceFromCrc
	ZoomMetrics["VideoDeviceFromCrc - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceFromCrc.AvgLoss)
	ZoomMetrics["VideoDeviceFromCrc - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceFromCrc.Bitrate)
	ZoomMetrics["VideoDeviceFromCrc - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceFromCrc.Jitter)
	ZoomMetrics["VideoDeviceFromCrc - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceFromCrc.Latency)
	ZoomMetrics["VideoDeviceFromCrc - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceFromCrc.MaxLoss)
	// VideoDeviceToCrc
	ZoomMetrics["VideoDeviceToCrc - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceToCrc.AvgLoss)
	ZoomMetrics["VideoDeviceToCrc - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceToCrc.Bitrate)
	ZoomMetrics["VideoDeviceToCrc - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceToCrc.Jitter)
	ZoomMetrics["VideoDeviceToCrc - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceToCrc.Latency)
	ZoomMetrics["VideoDeviceToCrc - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoDeviceToCrc.MaxLoss)
	//VideoInput
	ZoomMetrics["VideoInput - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoInput.AvgLoss)
	ZoomMetrics["VideoInput - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoInput.Bitrate)
	ZoomMetrics["VideoInput - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoInput.Jitter)
	ZoomMetrics["VideoInput - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoInput.Latency)
	ZoomMetrics["VideoInput - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoInput.MaxLoss)
	ZoomMetrics["VideoInput - Resolution"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoInput.Resolution)
	ZoomMetrics["VideoInput - FrameRate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoInput.FrameRate)
	//VideoOutput
	ZoomMetrics["VideoOutput - AvgLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoOutput.AvgLoss)
	ZoomMetrics["VideoOutput - Bitrate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoOutput.Bitrate)
	ZoomMetrics["VideoOutput - Jitter"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoOutput.Jitter)
	ZoomMetrics["VideoOutput - Latency"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoOutput.Latency)
	ZoomMetrics["VideoOutput - MaxLoss"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoOutput.MaxLoss)
	ZoomMetrics["VideoOutput - Resolution"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoOutput.Resolution)
	ZoomMetrics["VideoOutput - FrameRate"] = StringtoFloat(Meetingparticipantsclient.UserQos[0].VideoOutput.FrameRate)

	DashBoardMeetingClient, err := client.ListDashboardMeetings()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Below are metrics of Meeting from Dashboard API")

	ZoomMetrics["Meeting - Host"] = StringtoFloat(DashBoardMeetingClient.Meetings[0].Host)
	ZoomMetrics["Meeting - Duration"] = Durationsplit(DashBoardMeetingClient.Meetings[0].Duration)
	ZoomMetrics["Meeting - Has3RdPartyAudio"] = Booltofloat(DashBoardMeetingClient.Meetings[0].Has3RdPartyAudio)
	ZoomMetrics["Meeting - HasArchiving"] = Booltofloat(DashBoardMeetingClient.Meetings[0].HasArchiving)
	ZoomMetrics["Meeting - HasPstn"] = Booltofloat(DashBoardMeetingClient.Meetings[0].HasPstn)
	ZoomMetrics["Meeting - HasRecording"] = Booltofloat(DashBoardMeetingClient.Meetings[0].HasRecording)
	ZoomMetrics["Meeting - HasScreenShare"] = Booltofloat(DashBoardMeetingClient.Meetings[0].HasScreenShare)
	ZoomMetrics["Meeting - HasSip"] = Booltofloat(DashBoardMeetingClient.Meetings[0].HasSip)
	ZoomMetrics["Meeting - HasVideo"] = Booltofloat(DashBoardMeetingClient.Meetings[0].HasVideo)
	ZoomMetrics["Meeting - HasVoip"] = Booltofloat(DashBoardMeetingClient.Meetings[0].HasVoip)
	ZoomMetrics["Meeting - Participants"] = float64(DashBoardMeetingClient.Meetings[0].Participants)
	ZoomMetrics["Meeting - ID"] = float64(DashBoardMeetingClient.Meetings[0].ID)

	DashBoardWebinarClient, err := client.ListDashboardWebinar()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are metrics of Webinar from Dashboard API")

	ZoomMetrics["Webinar - Host"] = StringtoFloat(DashBoardWebinarClient.DashboardWebinar[0].Host)
	ZoomMetrics["Webinar - Duration"] = Durationsplit(DashBoardWebinarClient.DashboardWebinar[0].Duration)
	ZoomMetrics["Webinar - Has3RdPartyAudio"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].Has3RdPartyAudio)
	ZoomMetrics["Webinar - HasArchiving"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].HasArchiving)
	ZoomMetrics["Webinar - HasPstn"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].HasPstn)
	ZoomMetrics["Webinar - HasRecording"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].HasRecording)
	ZoomMetrics["Webinar - HasScreenShare"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].HasScreenShare)
	ZoomMetrics["Webinar - HasSip"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].HasSip)
	ZoomMetrics["Webinar - HasVideo"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].HasVideo)
	ZoomMetrics["Webinar - HasVoip"] = Booltofloat(DashBoardWebinarClient.DashboardWebinar[0].HasVoip)
	ZoomMetrics["Webinar - Participants"] = float64(DashBoardWebinarClient.DashboardWebinar[0].Participants)
	ZoomMetrics["Webinar - ID"] = float64(DashBoardWebinarClient.DashboardWebinar[0].ID)

	Webinarparticipantsclient, err := client.GetDashboardWebinarParticipants("")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are metrics of Webinar Participants QoS from Dashboard API")

	// AsDeviceFromCrc
	ZoomMetrics["Webinar_AsDeviceFromCrc - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceFromCrc.AvgLoss)
	ZoomMetrics["Webinar_AsDeviceFromCrc - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceFromCrc.Bitrate)
	ZoomMetrics["Webinar_AsDeviceFromCrc - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceFromCrc.Jitter)
	ZoomMetrics["Webinar_AsDeviceFromCrc - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceFromCrc.Latency)
	ZoomMetrics["Webinar_AsDeviceFromCrc - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceFromCrc.MaxLoss)
	//AsDeviceToCrc
	ZoomMetrics["Webinar_AsDeviceToCrc - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceToCrc.AvgLoss)
	ZoomMetrics["Webinar_AsDeviceToCrc - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceToCrc.Bitrate)
	ZoomMetrics["Webinar_AsDeviceToCrc - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceToCrc.Jitter)
	ZoomMetrics["Webinar_AsDeviceToCrc - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceToCrc.Latency)
	ZoomMetrics["Webinar_AsDeviceToCrc - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsDeviceToCrc.MaxLoss)
	//AsInput
	ZoomMetrics["Webinar_AsInput - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsInput.AvgLoss)
	ZoomMetrics["Webinar_AsInput - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsInput.Bitrate)
	ZoomMetrics["Webinar_AsInput - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsInput.Jitter)
	ZoomMetrics["Webinar_AsInput - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsInput.Latency)
	ZoomMetrics["Webinar_AsInput - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsInput.MaxLoss)
	ZoomMetrics["Webinar_AsInput - Resolution"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsInput.Resolution)
	ZoomMetrics["Webinar_AsInput - FrameRate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsInput.FrameRate)
	//AsOutput
	ZoomMetrics["Webinar_AsOutput - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.AvgLoss)
	ZoomMetrics["Webinar_AsOutput - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.Bitrate)
	ZoomMetrics["Webinar_AsOutput - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.Jitter)
	ZoomMetrics["Webinar_AsOutput - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.Latency)
	ZoomMetrics["Webinar_AsOutput - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.MaxLoss)
	ZoomMetrics["Webinar_AsOutput - Resolution"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.Resolution)
	ZoomMetrics["Webinar_AsOutput - FrameRate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AsOutput.FrameRate)
	// AudioDeviceFromCrc
	ZoomMetrics["Webinar_AudioDeviceFromCrc - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceFromCrc.AvgLoss)
	ZoomMetrics["Webinar_AudioDeviceFromCrc - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceFromCrc.Bitrate)
	ZoomMetrics["Webinar_AudioDeviceFromCrc - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceFromCrc.Jitter)
	ZoomMetrics["Webinar_AudioDeviceFromCrc - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceFromCrc.Latency)
	ZoomMetrics["Webinar_AudioDeviceFromCrc - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceFromCrc.MaxLoss)
	// AudioDeviceToCrc
	ZoomMetrics["Webinar_AudioDeviceToCrc - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceToCrc.AvgLoss)
	ZoomMetrics["Webinar_AudioDeviceToCrc - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceToCrc.Bitrate)
	ZoomMetrics["Webinar_AudioDeviceToCrc - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceToCrc.Jitter)
	ZoomMetrics["Webinar_AudioDeviceToCrc - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceToCrc.Latency)
	ZoomMetrics["Webinar_AudioDeviceToCrc - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioDeviceToCrc.MaxLoss)
	// AudioInput
	ZoomMetrics["Webinar_AudioInput - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioInput.AvgLoss)
	ZoomMetrics["Webinar_AudioInput - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioInput.Bitrate)
	ZoomMetrics["Webinar_AudioInput - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioInput.Jitter)
	ZoomMetrics["Webinar_AudioInput - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioInput.Latency)
	ZoomMetrics["Webinar_AudioInput - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioInput.MaxLoss)
	// AudioOutput
	ZoomMetrics["Webinar_AudioOutput - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioOutput.AvgLoss)
	ZoomMetrics["Webinar_AudioOutput - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioOutput.Bitrate)
	ZoomMetrics["Webinar_AudioOutput - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioOutput.Jitter)
	ZoomMetrics["Webinar_AudioOutput - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioOutput.Latency)
	ZoomMetrics["Webinar_AudioOutput - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_AudioOutput.MaxLoss)
	//CPUUsage
	ZoomMetrics["Webinar_CPUUsage - SystemMaxCPUUsage"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_CPUUsage.SystemMaxCPUUsage)
	ZoomMetrics["Webinar_CPUUsage - ZoomAvgCPUUsage"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_CPUUsage.ZoomAvgCPUUsage)
	ZoomMetrics["Webinar_CPUUsage - ZoomMaxCPUUsage"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_CPUUsage.ZoomMaxCPUUsage)
	ZoomMetrics["Webinar_CPUUsage - ZoomMinCPUUsage"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_CPUUsage.ZoomMinCPUUsage)

	// VideoDeviceFromCrc
	ZoomMetrics["Webinar_VideoDeviceFromCrc - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceFromCrc.AvgLoss)
	ZoomMetrics["Webinar_VideoDeviceFromCrc - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceFromCrc.Bitrate)
	ZoomMetrics["Webinar_VideoDeviceFromCrc - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceFromCrc.Jitter)
	ZoomMetrics["Webinar_VideoDeviceFromCrc - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceFromCrc.Latency)
	ZoomMetrics["Webinar_VideoDeviceFromCrc - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceFromCrc.MaxLoss)
	// VideoDeviceToCrc
	ZoomMetrics["Webinar_VideoDeviceToCrc - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceToCrc.AvgLoss)
	ZoomMetrics["Webinar_VideoDeviceToCrc - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceToCrc.Bitrate)
	ZoomMetrics["Webinar_VideoDeviceToCrc - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceToCrc.Jitter)
	ZoomMetrics["Webinar_VideoDeviceToCrc - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceToCrc.Latency)
	ZoomMetrics["Webinar_VideoDeviceToCrc - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoDeviceToCrc.MaxLoss)
	//VideoInput
	ZoomMetrics["Webinar_VideoInput - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.AvgLoss)
	ZoomMetrics["Webinar_VideoInput - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.Bitrate)
	ZoomMetrics["Webinar_VideoInput - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.Jitter)
	ZoomMetrics["Webinar_VideoInput - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.Latency)
	ZoomMetrics["Webinar_VideoInput - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.MaxLoss)
	ZoomMetrics["Webinar_VideoInput - Resolution"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.Resolution)
	ZoomMetrics["Webinar_VideoInput - FrameRate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoInput.FrameRate)
	//VideoOutput
	ZoomMetrics["Webinar_VideoOutput - AvgLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.AvgLoss)
	ZoomMetrics["Webinar_VideoOutput - Bitrate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.Bitrate)
	ZoomMetrics["Webinar_VideoOutput - Jitter"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.Jitter)
	ZoomMetrics["Webinar_VideoOutput - Latency"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.Latency)
	ZoomMetrics["Webinar_VideoOutput - MaxLoss"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.MaxLoss)
	ZoomMetrics["Webinar_VideoOutput - Resolution"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.Resolution)
	ZoomMetrics["Webinar_VideoOutput - FrameRate"] = StringtoFloat(Webinarparticipantsclient.Participants[0].Webinar_UserQos[0].Webinar_VideoOutput.FrameRate)

	// Anushree wrote the below
	Getchatmetrics, err := client.GetDashboardChat()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are the Chat Metrics from Dashboard API")
	ZoomMetrics["Page Size"] = float64(Getchatmetrics.PageSize)
	ZoomMetrics["Audio sent"] = float64(Getchatmetrics.Chats[0].AudioSent)
	ZoomMetrics["CodeSippet Sent"] = float64(Getchatmetrics.Chats[0].CodeSippetSent)
	ZoomMetrics["File sent"] = float64(Getchatmetrics.Chats[0].FilesSent)
	ZoomMetrics["Giphys sent"] = float64(Getchatmetrics.Chats[0].GiphysSent)
	ZoomMetrics["Group sent"] = float64(Getchatmetrics.Chats[0].GroupSent)
	ZoomMetrics["Image sent"] = float64(Getchatmetrics.Chats[0].ImagesSent)
	ZoomMetrics["P2P sent"] = float64(Getchatmetrics.Chats[0].P2PSent)
	ZoomMetrics["Text sent"] = float64(Getchatmetrics.Chats[0].TextSent)
	ZoomMetrics["Total sent"] = float64(Getchatmetrics.Chats[0].TotalSent)
	ZoomMetrics["UserID"] = StringtoFloat(Getchatmetrics.Chats[0].UserID)
	ZoomMetrics["Video sent"] = float64(Getchatmetrics.Chats[0].VideoSent)

	//DashboardSharingRecordingDetails
	GetSharingRecordingDetails, err := client.GetDashboardSharingRecordingDetails(" ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are the SharingRecording Metrics from Dashboard API")
	ZoomMetrics["UserID"] = StringtoFloat(GetSharingRecordingDetails.Participants[0].ID)
	ZoomMetrics["EndTime"] = StringtoFloat(GetSharingRecordingDetails.Participants[0].Details[0].EndTime)
	ZoomMetrics["StartTime"] = StringtoFloat(GetSharingRecordingDetails.Participants[0].Details[0].StartTime)

	//ReportBilling

	Getreportbilling, err := client.GetReportsBillings()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are the Billing Metrics from Reports API")
	ZoomMetrics["ID"] = StringtoFloat(Getreportbilling.BillingReport[0].ID)
	ZoomMetrics["TaxAmount"] = StringtoFloat(Getreportbilling.BillingReport[0].TaxAmount)
	ZoomMetrics["TotalAmount"] = StringtoFloat(Getreportbilling.BillingReport[0].TotalAmount)

	//ReportCloudRecording

	GetCloudRecording, err := client.GetReportsCloudRecording()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are the Cloud Recording Metrics from Reports API")
	ZoomMetrics["FreeUsage"] = StringtoFloat(GetCloudRecording.CloudRecordingStorage[0].FreeUsage)
	ZoomMetrics["PlanUsage"] = StringtoFloat(GetCloudRecording.CloudRecordingStorage[0].PlanUsage)
	ZoomMetrics["Usage"] = StringtoFloat(GetCloudRecording.CloudRecordingStorage[0].Usage)

	//planusage
	Getplanusagemetrics, err := client.GetBillingPlanUsage("")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are the PlanUsage Metrics from Billing API")

	ZoomMetrics["PlanBase - Hosts"] = float64(Getplanusagemetrics.PlanBase.Hosts)
	ZoomMetrics["PlanBase - Usage"] = float64(Getplanusagemetrics.PlanBase.Usage)
	ZoomMetrics["PlanBase - Pending"] = float64(Getplanusagemetrics.PlanBase.Pending)
	ZoomMetrics["PlanLargeMeeting - Hosts"] = float64(Getplanusagemetrics.PlanLargeMeeting[0].Hosts)
	ZoomMetrics["PlanLargeMeeting - Usage"] = float64(Getplanusagemetrics.PlanLargeMeeting[0].Usage)
	ZoomMetrics["PlanLargeMeeting - Pending"] = float64(Getplanusagemetrics.PlanLargeMeeting[0].Pending)
	ZoomMetrics["PlanUnited - Hosts"] = float64(Getplanusagemetrics.PlanUnited.Hosts)
	ZoomMetrics["PlanUnited - Usage"] = float64(Getplanusagemetrics.PlanUnited.Usage)
	ZoomMetrics["PlanUnited - Pending"] = float64(Getplanusagemetrics.PlanUnited.Pending)
	ZoomMetrics["PlanWebinar - Hosts"] = float64(Getplanusagemetrics.PlanWebinar[0].Hosts)
	ZoomMetrics["PlanWebinar - Hosts"] = float64(Getplanusagemetrics.PlanWebinar[0].Usage)
	ZoomMetrics["PlanWebinar - Hosts"] = float64(Getplanusagemetrics.PlanWebinar[0].Pending)
	ZoomMetrics["PlanZoomEvents - Hosts"] = float64(Getplanusagemetrics.PlanZoomEvents[0].Hosts)
	ZoomMetrics["PlanZoomEvents - Hosts"] = float64(Getplanusagemetrics.PlanZoomEvents[0].Usage)
	ZoomMetrics["PlanZoomEvents - Hosts"] = float64(Getplanusagemetrics.PlanZoomEvents[0].Pending)
	ZoomMetrics["PlanZoomRooms - Hosts"] = float64(Getplanusagemetrics.PlanZoomRooms.Hosts)
	ZoomMetrics["PlanZoomRooms - Hosts"] = float64(Getplanusagemetrics.PlanZoomRooms.Usage)

	//reportsmeetings
	Getreportsmeetingsmetrics, err := client.GetReportsMeetings(" ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are the Meetings Metrics from Reports  API")
	ZoomMetrics["RepMeetings - Duration"] = float64(Getreportsmeetingsmetrics.RepMeetings[0].Duration)
	ZoomMetrics["RepMeetings - ID"] = float64(Getreportsmeetingsmetrics.RepMeetings[0].ID)
	ZoomMetrics["RepMeetings - ParticipantsCount"] = float64(Getreportsmeetingsmetrics.RepMeetings[0].ParticipantsCount)
	ZoomMetrics["RepMeetings - TotalMinutes"] = float64(Getreportsmeetingsmetrics.RepMeetings[0].TotalMinutes)
	ZoomMetrics["RepMeetings - Type"] = float64(Getreportsmeetingsmetrics.RepMeetings[0].Type)

	//webinardetails
	GetReportsWebinarDetailsMetrics, err := client.GetReportsWebinarDetails("")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Below are the Webinar Metrics from Reports API")

	ZoomMetrics["GetReportsWebinarDetailsMetrics - Duration"] = float64(GetReportsWebinarDetailsMetrics.Duration)
	ZoomMetrics["GetReportsWebinarDetailsMetrics - ID"] = StringtoFloat(GetReportsWebinarDetailsMetrics.ID)
	ZoomMetrics["GetReportsWebinarDetailsMetrics - Topic"] = StringtoFloat(GetReportsWebinarDetailsMetrics.Topic)
	ZoomMetrics["GetReportsWebinarDetailsMetrics - TotalTime"] = float64(GetReportsWebinarDetailsMetrics.TotalTime)
	ZoomMetrics["GetReportsWebinarDetailsMetrics - Type"] = float64(GetReportsWebinarDetailsMetrics.Type)

	//reportdailyusage
	GetReportDailyUsageMetrics, err := client.GetReportsDailyUsage()
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Below are the Daily Usage Metrics from Report  API")
	ZoomMetrics["GetReportDailyUsageMetrics - MeetingMinutes"] = StringtoFloat(GetReportDailyUsageMetrics.Dates[0].MeetingMinutes)
	ZoomMetrics["GetReportDailyUsageMetrics - Meetings"] = StringtoFloat(GetReportDailyUsageMetrics.Dates[0].Meetings)
	ZoomMetrics["GetReportDailyUsageMetrics - NewUsers"] = StringtoFloat(GetReportDailyUsageMetrics.Dates[0].NewUsers)
	ZoomMetrics["GetReportDailyUsageMetrics - Participants"] = StringtoFloat(GetReportDailyUsageMetrics.Dates[0].Participants)

	return ZoomMetrics, err
}
