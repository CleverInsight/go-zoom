/* Author : Bastin Robin
Created on : 5th April 2022
Description : This is a Module to Get the users from the API,

*/
package zoom

import (
	"encoding/json"
	"log"
)

type UserResponse struct {
	PageCount    int32 `json:"page_count"`
	PageNumber   int32 `json:"page_number"`
	PageSize     int32 `json:"page_size"`
	TotalRecords int32 `json:"total_records"`
	Users        Users
}

type Users []User

type User struct {
	Id        string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string
	Type      int32
	Pmi       int64
	Verified  int16
	RoleId    string
	Avatar    string `json:"pic_url"`
}

func (z *Zoom) GetUsers() (UserResponse, error) {
	response, err := z.ReqBody("GET", "/users")
	if err != nil {
		log.Printf("Error during GetUsers %v", err)
		return UserResponse{}, err
	}

	// Unmarshal the response which is in byte format
	var userResponse UserResponse
	err = json.Unmarshal(response, &userResponse)
	if err != nil {
		log.Printf("Error during unmarshaling %v", err)
		return UserResponse{}, err
	}

	return userResponse, nil

}
