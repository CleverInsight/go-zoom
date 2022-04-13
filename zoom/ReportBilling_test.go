package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetReportBillings(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
            "billing_reports": [
              {
                "end_date": "",
                "id": "USR111",
                "start_date": "",
                "tax_amount": "",
                "total_amount": "",
                "type": ""
              }
            ],
            "currency": "",
            "download_url": ""
          }`))
	}))
	client := &Zoom{
		Endpoint: testServer.URL,
	}
	billingreport, err := client.GetBillings()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, billingreport, " results was empty")
	assert.Equal(t, billingreport.BillingReport[0].ID, "USR111", "ID not been matched")
	assert.Equal(t, billingreport.BillingReport[0].TaxAmount, "", "TaxAmount not been matched")
	assert.Equal(t, billingreport.BillingReport[0].TotalAmount, "", "TotalAmount not been matched")
}
