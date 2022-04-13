package zoom

import (
	"encoding/json"
	"log"
)

type Billings struct {
	Currency      string        `json:"currency"`
	DownloadURL   string        `json:"download_url"`
	BillingReport BillingReport `json:"billing_reports"`
}
type BillingReport []BillingReports
type BillingReports struct {
	EndDate     string `json:"end_date"`
	ID          string `json:"id"` //Metric
	StartDate   string `json:"start_date"`
	TaxAmount   string `json:"tax_amount"`   //Metric
	TotalAmount string `json:"total_amount"` //Metric
	Type        string `json:"type"`
}

func (z *Zoom) GetBillings() (Billings, error) {

	billings, err := z.ReqBody("GET", "/report/billing")
	if err != nil {
		log.Println(err)
		return Billings{}, err
	}

	var billing Billings

	// Unmarshal the response into Response struct
	err = json.Unmarshal(billings, &billing)
	if err != nil {
		log.Println(err)
		return Billings{}, err
	}
	return billing, nil
}
