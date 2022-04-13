package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDashboardChat(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
      "from": "",
      "next_page_token": "",
      "page_size": ,
      "to": "",
      "users": [
        {
          "audio_sent": 20,
          "code_sippet_sent": 20,
          "email": "",
          "files_sent": 20,
          "giphys_sent": 20,
          "group_sent": 20,
          "images_sent": 20,
          "p2p_sent": 20,
          "text_sent": 20,
          "total_sent": 20,
          "user_id": "",
          "user_name": "aaa",
          "video_sent": 20
        }
      ]
    }`))
	}))
	client := &Zoom{
		Endpoint: testServer.URL,
	}
	chat, err := client.GetDashboardChat()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, chat, " response was empty")
	assert.Equal(t, chat.Chats[0].AudioSent, 20, "AudioSent not been matched")
	assert.Equal(t, chat.Chats[0].CodeSippetSent, 20, "CodeSippetSent not been matched")
	assert.Equal(t, chat.Chats[0].FilesSent, 20, "FileSent not been matched")
	assert.Equal(t, chat.Chats[0].GiphysSent, 20, "GiphysSent not been matched")
	assert.Equal(t, chat.Chats[0].GroupSent, 20, "GroupSent not been matched")
	assert.Equal(t, chat.Chats[0].ImagesSent, 20, "ImageSent not been matched")
	assert.Equal(t, chat.Chats[0].P2PSent, 20, "P2PSent not been matched")
	assert.Equal(t, chat.Chats[0].TextSent, 20, "TextSent not been matched")
	assert.Equal(t, chat.Chats[0].TotalSent, 20, "TotalSent not been matched")
	assert.Equal(t, chat.Chats[0].UserName, "aaa", "FileSent not been matched")
	assert.Equal(t, chat.Chats[0].VideoSent, 20, "VideoSent not been matched")
}
