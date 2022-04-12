package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChat(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
            "from": "",
            "next_page_token": "",
            "page_size": "",
            "to": "",
            "users": [
              {
                "audio_sent": ,
                "code_sippet_sent": ,
                "email": "",
                "files_sent": ,
                "giphys_sent": ,
                "group_sent": ,
                "images_sent": ,
                "p2p_sent": ,
                "text_sent": ,
                "total_sent": ,
                "user_id": "",
                "user_name": "",
                "video_sent": 
              }
            ]
          }`))
	}))
	client := &Zoom{
		Endpoint: testServer.URL,
	}
	chat, err := client.GetChat()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, chat, " results was empty")
	assert.Equal(t, chat.PageSize, "", "PageSize is not matching")
}
