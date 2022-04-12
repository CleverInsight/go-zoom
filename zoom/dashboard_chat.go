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
	AudioSent      int    `json:"audio_sent"`       //Metric
	CodeSippetSent int    `json:"code_sippet_sent"` //Metric
	Email          string `json:"email"`
	FilesSent      int    `json:"files_sent"`  //Metric
	GiphysSent     int    `json:"giphys_sent"` //Metric
	GroupSent      int    `json:"group_sent"`  //Metric
	ImagesSent     int    `json:"images_sent"` //Metric
	P2PSent        int    `json:"p2p_sent"`    //Metric
	TextSent       int    `json:"text_sent"`   //Metric
	TotalSent      int    `json:"total_sent"`  //Metric
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`  //Metric
	VideoSent      int    `json:"video_sent"` //Metric
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
