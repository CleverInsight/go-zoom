package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_count": 1,
			"page_number": 1,
			"page_size": 30,
			"total_records": 1,
			"next_page_token": "",
			"users": [
			  {
				"id": "WytA-G2_SICnqEjr988zcA",
				"first_name": "Bastin",
				"last_name": "Robins J",
				"email": "bastinrobins@gmail.com",
				"type": 1,
				"pmi": 8722824397,
				"timezone": "Asia/Kolkata",
				"verified": 1,
				"dept": "",
				"created_at": "2020-09-26T03:48:43Z",
				"last_login_time": "2022-04-05T10:31:30Z",
				"last_client_version": "5.9.1.3506(mac)",
				"pic_url": "https://lh3.googleusercontent.com/a-/AOh14GhNvcwafxOv7kxZMDQPfnDSaOjdIOwqbPNcpQm9Dw=s96-c",
				"language": "en-US",
				"phone_number": "",
				"status": "active",
				"role_id": "0"
			  }
			]
		  }`))
	}))

	client := &Zoom{
		Endpoint: testServer.URL,
	}

	users, err := client.GetUsers()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, users, "Users response was empty")
	assert.Equal(t, users.PageCount, int32(1), "Count has not been matched")
}
