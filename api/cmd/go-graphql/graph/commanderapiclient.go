package graph

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/go-graphql/graph/model"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/pb/admissions"
	uid "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/pb/type/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	// cmdr "gitlab.com/2ndwatch/microservices/apis/ms-api-commander/pkg/pb/commander"
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
	targetApiEndpoint = getenv("TARGET_API_ENDPOINT", "<TBD>/commands")
)

func getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

//region private funcs
func (apiClient *CommanderClient) fetchXApiKeyHeaders() string {
	// TODO: once auth endpoint is available from API, revisit this and implement auth header get
	return "foo"
}

func convertAdmissionPolicyToProtoStruct(message model.AdmissionPolicy) structpb.Struct {
	principals := []string{}
	for _, val := range message.Principal {
		if val != nil {
			principals = append(principals, *val)
		}
	}

	actions := []string{}
	for _, val := range message.Actions {
		if val != nil {
			actions = append(actions, *val)
		}
	}

	resources := []string{}
	for _, val := range message.Resources {
		if val != nil {
			resources = append(resources, *val)
		}
	}
	effect := admissions.Effect(admissions.Effect_value[message.Effect.String()])
	pbAdmissionsMessage := &admissions.AdmissionMessage{
		Id:         &uid.UUID{Value: message.ID},
		Name:       message.Name,
		Effect:     effect,
		Type:       admissions.AdmissionPolicyType(admissions.AdmissionPolicyType_value[message.Type.String()]),
		Principals: principals,
		Actions:    actions,
		Resources:  resources,
	}

	msg := &structpb.Struct{}
	protojson.Unmarshal([]byte(pbAdmissionsMessage.String()), msg)
	return *msg
}

//endregion private funcs

//region public funcs
func (apiClient *CommanderClient) MakeApiRequest(message model.AdmissionPolicy, action string, requestType string) (*string, error) {

	// requestBody, err := json.Marshal(message)
	requestBody := convertAdmissionPolicyToProtoStruct(message)
	commandParams := &admissions.CommandParams{
		Action: action,
		Data:   &requestBody,
	}

	if requestType == "" {
		return nil, fmt.Errorf("must specific request type (POST, PUT, DELETE)")
	}
	// TODO: test if below will properly represent commandParams when sending to API -- seems unlikely
	byte_slice := []byte{}
	protojson.Unmarshal(byte_slice, commandParams)
	request, err := http.NewRequest(requestType, targetApiEndpoint, bytes.NewBuffer(byte_slice))

	if err != nil {
		return nil, err
	}

	// request.Header.Set("x-api-key", apiClient.fetchXApiKeyHeaders())
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("encountered an error while attempting to request to API %s, %e", targetApiEndpoint, err)
	}

	if response.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("encountered an error while reading API response from API request %s, %e", targetApiEndpoint, err)
		}

		//Convert bytes to String and print
		jsonStr := string(body)
		return &jsonStr, nil

	} else {
		return nil, fmt.Errorf("API request failed with error: %v", response.Status)
	}
}

//endregion public funcs