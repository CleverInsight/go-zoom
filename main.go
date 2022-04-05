package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cleverinsight/go-zoom/zoom"
)

func main() {

	headers := make(map[string][]string)
	token := fmt.Sprintf("Bearer %v", os.Getenv("ZOOM_BEARER"))
	headers["Authorization"] = []string{token}

	client := zoom.Zoom{
		ApiKey:   os.Getenv("ZOOM_API_KEY"),
		ApiToken: os.Getenv("ZOOM_API_TOKEN"),
		Headers:  headers,
		Endpoint: "https://api.zoom.us/v2",
	}

	users, err := client.GetUsers()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(users.PageCount)
	fmt.Println(users.PageSize)
	fmt.Println(users.TotalRecords)

	for _, user := range users.Users {
		fmt.Println(user.FirstName, user.LastName)
	}
}
