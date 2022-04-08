package zoom

import (
	"encoding/json"
	"log"
)

type Billings struct {
	BillingReports BillingReports
	Currency       string `json:"currency"`
	DownloadURL    string `json:"download_url"`
}

type BillingReports []struct {
	EndDate   string `json:"end_date"`
	ID        string `json:"id"`
	StartDate string `json:"start_date"`
	TaxAmt    string `json:"tax_amount"`   //Metric
	TotalAmt  string `json:"total_amount"` // Metric
	Type      string `json:"type"`
}

func (z *Zoom) GetBillings() (Billings, error) {
	billings, err := z.ReqBody("/report/billing", "GET")
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
