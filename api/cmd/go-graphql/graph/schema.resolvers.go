package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/go-graphql/graph/generated"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/go-graphql/graph/model"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/pb/admissions"
	uid "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/pb/type/uuid"
)

// CreateAdmissionPolicy is the resolver for the createAdmissionPolicy field.
func (r *mutationResolver) CreateAdmissionPolicy(ctx context.Context, admissionPolicy model.AdmissionPolicyInput) (*model.AdmissionPolicy, error) {
	policyUuid := strings.Replace(uuid.New().String(), "-", "", -1)

	createdAdmissionPolicy := &model.AdmissionPolicy{
		ID:        policyUuid,
		Name:      admissionPolicy.Name,
		Effect:    admissionPolicy.Effect,
		Type:      admissionPolicy.Type,
		Principal: append([]*string{}, admissionPolicy.Principal...),
		Actions:   append([]*string{}, admissionPolicy.Actions...),
		Resources: append([]*string{}, admissionPolicy.Resources...),
	}
	// // TODO: Send marshalled JSON object to Commander API for database inserts here
	r.admissionPolicies = append(r.admissionPolicies, createdAdmissionPolicy)

	// encode input struct
	encodedBlob, err := EncodeToString(createdAdmissionPolicy)
	if err != nil {
		return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	}
	messageData := CommanderMessage{
		Action: "CreateAdmissionPolicy",
		Data: CommanderMessageData{
			MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
			DataBlob:    encodedBlob,
			CommandName: "CreateAdmissionPolicy",
		},
	}

	uuid := &uid.UUID{Value: policyUuid}
	effect := admissions.Effect(admissions.Effect_value[admissionPolicy.Effect.String()])
	policyType := admissions.AdmissionPolicyType(admissions.AdmissionPolicyType_value[admissionPolicy.Type.String()])
	princips := []string{}
	for _, val := range admissionPolicy.Principal {
		if val != nil {
			princips = append(princips, *val)
		}
	}

	acts := []string{}
	for _, val := range admissionPolicy.Actions {
		if val != nil {
			acts = append(acts, *val)
		}
	}

	ress := []string{}
	for _, val := range admissionPolicy.Resources {
		if val != nil {
			ress = append(ress, *val)
		}
	}

	pbAdmissionsMessage := &admissions.AdmissionMessage{
		Id:         uuid,
		Name:       admissionPolicy.Name,
		Effect:     effect,
		Type:       policyType,
		Principals: princips,
		Actions:    acts,
		Resources:  ress,
	}

	// TODO: marshal above pbAdmissionsMessage to proto struct for sending to Commander API in below func call
	// commandParams := &admissions.CommandParams{
	// 	Action: "CreateAdmissionPolicy",
	// 	Data:   protojson.Marshal(),
	// }
	response, err := r.apiClient.MakeApiRequest(messageData, "POST")
	if err != nil {
		return nil, fmt.Errorf("encountered an error while trying to POST object: %v", err)
	}

	fmt.Printf("Successfully submitted POST request for object %s", *response)
	return createdAdmissionPolicy, nil
}

// UpdateAdmissionPolicyActions is the resolver for the updateAdmissionPolicyActions field.
func (r *mutationResolver) UpdateAdmissionPolicyActions(ctx context.Context, id string, admissionPolicyActions []*string) (*model.AdmissionPolicy, error) {
	updateAdmissionPolicyActionsModel := &model.AdmissionPolicyActions{
		ID:      id,
		Actions: admissionPolicyActions,
	}
	// encode input struct
	encodedBlob, err := EncodeToString(updateAdmissionPolicyActionsModel)
	if err != nil {
		return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	}

	messageData := CommanderMessage{
		Action: "UpdateAdmissionPolicyActions",
		Data: CommanderMessageData{
			MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
			DataBlob:    encodedBlob,
			CommandName: "UpdateAdmissionPolicyActions", // may not need this
		},
	}
	response, err := r.apiClient.MakeApiRequest(messageData, "PUT")
	if err != nil {
		return nil, fmt.Errorf("encountered an error while trying to PUT object: %v", err)
	}

	fmt.Printf("Successfully submitted PUT request for object %s", *response)
	// return updateAdmissionPolicyActionsModel, nil
	return &model.AdmissionPolicy{}, nil
}

// UpdateAdmissionPolicyPrincipals is the resolver for the updateAdmissionPolicyPrincipals field.
func (r *mutationResolver) UpdateAdmissionPolicyPrincipals(ctx context.Context, id string, admissionPolicyPrincipals []*string) (*model.AdmissionPolicy, error) {
	updateAdmissionPolicyPrincipalsModel := &model.AdmissionPolicyPrincipals{
		ID:         id,
		Principals: admissionPolicyPrincipals,
	}

	// encode input struct
	encodedBlob, err := EncodeToString(updateAdmissionPolicyPrincipalsModel)
	if err != nil {
		return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	}

	messageData := CommanderMessage{
		Action: "UpdateAdmissionPolicyPrincipals",
		Data: CommanderMessageData{
			MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
			DataBlob:    encodedBlob,
			CommandName: "UpdateAdmissionPolicyPrincipals", // may not need this
		},
	}
	response, err := r.apiClient.MakeApiRequest(messageData, "PUT")
	if err != nil {
		return nil, fmt.Errorf("encountered an error while trying to PUT object: %v", err)
	}

	fmt.Printf("Successfully submitted PUT request for object %s", *response)
	// return updateAdmissionPolicyPrincipalsModel, nil
	return &model.AdmissionPolicy{}, nil
}

// UpdateAdmissionPolicyResources is the resolver for the updateAdmissionPolicyResources field.
func (r *mutationResolver) UpdateAdmissionPolicyResources(ctx context.Context, id string, admissionPolicyResources []*string) (*model.AdmissionPolicy, error) {
	updateAdmissionPolicyResourcesModel := &model.AdmissionPolicyResources{
		ID:        id,
		Resources: admissionPolicyResources,
	}

	// encode input struct
	encodedBlob, err := EncodeToString(updateAdmissionPolicyResourcesModel)
	if err != nil {
		return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	}

	messageData := CommanderMessage{
		Action: "UpdateAdmissionPolicyResources",
		Data: CommanderMessageData{
			MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
			DataBlob:    encodedBlob,
			CommandName: "UpdateAdmissionPolicyResources", // may not need this
		},
	}
	response, err := r.apiClient.MakeApiRequest(messageData, "PUT")
	if err != nil {
		return nil, fmt.Errorf("encountered an error while trying to PUT object: %v", err)
	}

	fmt.Printf("Successfully submitted PUT request for object %s", *response)
	// return updateAdmissionPolicyResourcesModel, nil
	return &model.AdmissionPolicy{}, nil
}

// DeleteAdmissionPolicy is the resolver for the deleteAdmissionPolicy field.
func (r *mutationResolver) DeleteAdmissionPolicy(ctx context.Context, id string) (*bool, error) {
	// The below logic will be replaced with Commander API call for deletes here, this is temporary for example
	// TODO: Send marshalled JSON object to Commander API for database deletes below here
	var deleted = false
	messageData := CommanderMessage{
		Action: "DeleteAdmissionPolicy",
		Data: CommanderMessageData{
			MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
			DataBlob:    id,
			CommandName: "DeleteAdmissionPolicy", // may not need this
		},
	}
	response, err := r.apiClient.MakeApiRequest(messageData, "DELETE")
	if err != nil {
		return nil, fmt.Errorf("encountered an error while trying to DELETE object: %v", err)
	}
	deleted = true

	fmt.Printf("Successfully submitted DELETE request for object %s", *response)
	return &deleted, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todoUuid := strings.Replace(uuid.New().String(), "-", "", -1)
	userUuid := strings.Replace(uuid.New().String(), "-", "", -1)

	todo := &model.Todo{
		Text: input.Text,
		ID:   todoUuid,
		User: &model.User{ID: userUuid, Name: fmt.Sprintf("user '%s'", input.UserID)},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

// AdmissionPolicies is the resolver for the admissionPolicies field.
func (r *queryResolver) AdmissionPolicies(ctx context.Context, principal string, policyType *model.AdmissionPolicyType, policyName *string) ([]*model.AdmissionPolicy, error) {
	if principal != "" { // filter for admission policies w/ principals matching the provided principal string
		ret := []*model.AdmissionPolicy{}

		for index := range r.admissionPolicies {
			for xindex := range r.admissionPolicies[index].Principal {
				if *r.admissionPolicies[index].Principal[xindex] == principal {
					if policyName == nil && policyType == nil { // no additional filtering needed
						ret = append(ret, r.admissionPolicies[index])
					} else {
						if policyName != nil && policyType == nil { // filter for matching policyName only
							if r.admissionPolicies[index].Name == *policyName {
								ret = append(ret, r.admissionPolicies[index])
							}
						} else if policyType != nil && policyName == nil { // filter for matching policyType only
							if *r.admissionPolicies[index].Type == *policyType {
								ret = append(ret, r.admissionPolicies[index])
							}
						} else { // filter for matching policyName and matching policyType
							if r.admissionPolicies[index].Name == *policyName && *r.admissionPolicies[index].Type == *policyType {
								ret = append(ret, r.admissionPolicies[index])
							}
						}
					}
				}
			}
		}
		return ret, nil
	}
	return r.admissionPolicies, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
