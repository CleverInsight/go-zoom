/* Author : Bastin Robin
Created on : 5th April 2022
Description : This is a  the main module containing the Zoom Client,

*/
package zoom

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Zoom struct {
	ApiKey   string
	ApiToken string
	Endpoint string // https://api.zoom.com/v1/
	Headers  map[string][]string
}

// Getting updated bearer token
func (z *Zoom) GetToken() (string, error) {

	token := "Bearer "
	return token, nil
}

// Sending Http Request
func (z *Zoom) ReqBody(method string, url string) ([]byte, error) {

	endpoint := fmt.Sprintf("%v%v", z.Endpoint, url)

	// Create request object to make http calls
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Printf("http request failed %v", err)
		return []byte(""), err
	}

	// Passing the auth headers into given request client
	req.Header = z.Headers

	// Create a client to execute the above request
	client := http.Client{}

	// Execute the network call
	response, err := client.Do(req)
	if err != nil {
		log.Printf("Issue during network call %v", err)
		return []byte(""), err
	}

	// Convert the io buffer into bytes
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Issue during IO utility %v", err)
		return []byte(""), err
	}

	return data, nil

}
