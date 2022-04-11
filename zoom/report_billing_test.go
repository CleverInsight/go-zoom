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
                "id": "",
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
	//assert.Equal(t,meetingparticipants., "", "ID has not been matched")
}
