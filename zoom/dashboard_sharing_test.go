package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSharingRecordingDetails(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
            "from": "",
            "next_page_token": "",
            "page_size": "",
            "to": "",
            "users": [
              {
                "audio_sent": "",
                "code_sippet_sent": "",
                "email": "",
                "files_sent": "",
                "giphys_sent": "",
                "group_sent": "",
                "images_sent": "",
                "p2p_sent": "",
                "text_sent": "",
                "total_sent": "",
                "user_id": "",
                "user_name": "",
                "video_sent": ""
              }
            ]
          }`))
	}))
	client := &Zoom{
		Endpoint: testServer.URL,
	}
	dashboard_sharing, err := client.GetSharingRecordingDetails("")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, dashboard_sharing, " response was empty")
	// 	assert.Equal(t, dashboard_sharing.PageCount, 0, "PageCount not been matched")
	// 	assert.Equal(t, dashboard_sharing.PageSize, 0, "Pagesize not been matched")
}
