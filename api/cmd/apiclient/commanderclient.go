package apiclient

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CommanderClient struct {
}

type CommanderMessage struct {
	Action string               `json:"action"`
	Data   CommanderMessageData `json:"data"`
}

type CommanderMessageData struct {
	MessageUUID string `json:"message_uuid"`
	DataBlob    string `json:"data_blob"`    // event action value
	CommandName string `json:"command_name"` // event action
}

var (
	targetApiEndpoint = flag.String("TARGET_API_ENDPOINT", "<TBD>/commands", "The Commander API endpoint to make requests to")
)

//region private funcs
func (apiClient *CommanderClient) fetchXApiKeyHeaders() string {
	// TODO: once auth endpoint is available from API, revisit this and implement auth header get
	return "foo"
}

//endregion private funcs

//region public funcs
func (apiClient *CommanderClient) MakeApiRequest(message CommanderMessage, requestType String) (*string, error) {

	requestBody, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	if requestType == nil || requestType == "" {
		return nil, fmt.Errorf("must specific request type (POST, PUT, DELETE)")
	}

	request, err := http.NewRequest(requestType, *targetApiEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	// request.Header.Set("x-api-key", apiClient.fetchXApiKeyHeaders())
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("encountered an error while attempting to request to API %s, %e", *targetApiEndpoint, err)
	}

	if response.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("encountered an error while reading API response from API request %s, %e", *targetApiEndpoint, err)
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		return &jsonStr, nil

	} else {
		return nil, fmt.Errorf("API request failed with error: %v", response.Status)
	}
}

//endregion public funcs
