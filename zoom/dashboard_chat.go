package zoom

import (
	"encoding/json"
	"log"
)

type Responses struct {
	From          string `json:"from"`
	NextPageToken string `json:"next_page_token"`
	PageSize      string `json:"page_size"`
	To            string `json:"to"`
	Chats         Chats
}
type Chats []Chat
type Chat struct {
	AudioSent      string `json:"audio_sent"`
	CodeSippetSent string `json:"code_sippet_sent"`
	Email          string `json:"email"`
	FilesSent      string `json:"files_sent"`
	GiphysSent     string `json:"giphys_sent"`
	GroupSent      string `json:"group_sent"`
	ImagesSent     string `json:"images_sent"`
	P2PSent        string `json:"p2p_sent"`
	TextSent       string `json:"text_sent"`
	TotalSent      string `json:"total_sent"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	VideoSent      string `json:"video_sent"`
}

// Return the list of all chats in zoom account
func (z *Zoom) GetChat() (Responses, error) {
	// chats, err := z.ReqBody("/metrics/chat", "GET")

	response, err := z.ReqBody("GET", "/metrics/chat")
	if err != nil {
		log.Println(err)
		return Responses{}, err
	}

	var chat Responses

	// Unmarshal the response into Response struct
	err = json.Unmarshal(response, &chat)
	if err != nil {
		log.Println(err)
		return Responses{}, err
	}
	return chat, nil
}
