package zoom

import (
	"encoding/json"
	"log"
)

type CloudRecordings struct {
	From                  string                 `json:"from"`
	To                    string                 `json:"to"`
	CloudRecordingStorage CloudRecordingStorages `json:"cloud_recording_storage"`
}
type CloudRecordingStorages []CloudRecordingStorage
type CloudRecordingStorage struct {
	Date      string `json:"date"`
	FreeUsage string `json:"free_usage"` //Metric
	PlanUsage string `json:"plan_usage"` //Metric
	Usage     string `json:"usage"`      //Metric
}

func (z *Zoom) GetReportCloudRecording() (CloudRecordings, error) {

	cloudrecordings, err := z.ReqBody("GET", "/report/cloud_recording")
	if err != nil {
		log.Println("Error in GetCloudRecording's Response", err)
		return CloudRecordings{}, err
	}

	var cloudrecord CloudRecordings

	err = json.Unmarshal(cloudrecordings, &cloudrecord)
	if err != nil {
		log.Println("Error in Unmarshalling response cloudrecordings", err)
		return CloudRecordings{}, err
	}

	return cloudrecord, err

}
