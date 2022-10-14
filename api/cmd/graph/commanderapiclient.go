package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/graph/model"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/pb/admissions"
	uid "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/pb/type/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

type CommanderClient struct {
}

type CommanderMessageData struct {
	MessageUUID string `json:"message_uuid"`
	DataBlob    string `json:"data_blob"`    // event action value
	CommandName string `json:"command_name"` // event action
}

var (
	targetApiEndpoint = getenv("TARGET_API_ENDPOINT", "http://127.0.0.1:8080/commands")
)

func getenv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

//region private funcs
func convertAdmissionPolicyToProtoStruct(message model.AdmissionPolicy) (*structpb.Struct, error) {
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
		Id:         &uid.UUID{Value: *message.ID},
		Name:       message.Name,
		Effect:     effect,
		Type:       admissions.AdmissionPolicyType(admissions.AdmissionPolicyType_value[message.Type.String()]),
		Principals: principals,
		Actions:    actions,
		Resources:  resources,
	}

	msg := &structpb.Struct{}
	marshalledMessage, err := json.Marshal(pbAdmissionsMessage)
	if err != nil {
		return nil, err
	}
	err = protojson.Unmarshal(marshalledMessage, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

//endregion private funcs

//region public funcs
func (apiClient *CommanderClient) MakeApiRequest(message model.AdmissionPolicy, action string) (*string, error) {

	// requestBody, err := json.Marshal(message)
	requestBody, err := convertAdmissionPolicyToProtoStruct(message)
	if err != nil {
		return nil, err
	}
	commandParams := &admissions.CommandParams{
		Action: action,
		Data:   requestBody,
		Sync:   true, // true for synchronous request
	}

	// TODO: test if below will properly represent commandParams when sending to API -- seems unlikely
	marshalledCommandParams, err := json.Marshal(commandParams)
	if err != nil {
		return nil, err
	}
	msg := &structpb.Struct{}
	err = protojson.Unmarshal(marshalledCommandParams, msg)
	if err != nil {
		return nil, err
	}
	// static value of POST for below request method param as all Commander API submissions are treated as message create ops
	request, err := http.NewRequest("POST", targetApiEndpoint, bytes.NewBuffer(marshalledCommandParams))
	if err != nil {
		return nil, err
	}

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
