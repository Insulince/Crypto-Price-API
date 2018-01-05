package twilio

import (
	"net/url"
	"net/http"
	"strings"
	"encoding/json"
	"fmt"
)

func m(config Config) () {
	message := "swag"

	bodyData := url.Values{}
	bodyData.Set("To", config.ReceiverPhoneNumber)
	bodyData.Set("From", config.SenderPhoneNumber)
	bodyData.Set("Body", message)
	httpClient := &http.Client{}
	request, err := http.NewRequest("POST", config.TwilioAPIURL, strings.NewReader(bodyData.Encode()))
	if err != nil {
		panic(err)
	}
	request.SetBasicAuth(config.TwilioAccountSID, config.TwilioAuthToken)
	request.Header.Add("Accept", "application/json")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := httpClient.Do(request)
	if err != nil {
		panic(err)
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		var twilioResponse TwilioResponse
		err := json.NewDecoder(response.Body).Decode(&twilioResponse)
		if err != nil {
			panic(err)
		}
		fmt.Println(twilioResponse)
	} else {
		fmt.Println(response.Status)
	}
}