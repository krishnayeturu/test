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
)

// CreateAdmissionPolicy is the resolver for the createAdmissionPolicy field.
func (r *mutationResolver) CreateAdmissionPolicy(ctx context.Context, admissionPolicy model.AdmissionPolicyInput) (*model.AdmissionPolicy, error) {
	policyUuid := strings.Replace(uuid.New().String(), "-", "", -1)

	createdAdmissionPolicy := &model.AdmissionPolicy{
		ID:        &policyUuid,
		Name:      admissionPolicy.Name,
		Effect:    admissionPolicy.Effect,
		Type:      admissionPolicy.Type,
		Principal: append([]*string{}, admissionPolicy.Principal...),
		Actions:   append([]*string{}, admissionPolicy.Actions...),
		Resources: append([]*string{}, admissionPolicy.Resources...),
	}
	// // TODO: Send marshalled JSON object to Commander API for database inserts here
	r.admissionPolicies = append(r.admissionPolicies, createdAdmissionPolicy)

	// region Database Operations
	insert_id, err := createdAdmissionPolicy.Insert()
	if err != nil {
		return nil, err
	}

	fmt.Println(insert_id)
	// enddregion Database Operations

	// encode input struct
	// encodedBlob, err := EncodeToString(createdAdmissionPolicy)
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	// }
	// messageData := CommanderMessage{
	// 	Action: "CreateAdmissionPolicy",
	// 	Data: CommanderMessageData{
	// 		MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
	// 		DataBlob:    encodedBlob,
	// 		CommandName: "CreateAdmissionPolicy",
	// 	},
	// }

	// uuid := &uid.UUID{Value: policyUuid}
	// effect := admissions.Effect(admissions.Effect_value[admissionPolicy.Effect.String()])
	// policyType := admissions.AdmissionPolicyType(admissions.AdmissionPolicyType_value[admissionPolicy.Type.String()])
	// princips := []string{}
	// for _, val := range admissionPolicy.Principal {
	// 	if val != nil {
	// 		princips = append(princips, *val)
	// 	}
	// }

	// acts := []string{}
	// for _, val := range admissionPolicy.Actions {
	// 	if val != nil {
	// 		acts = append(acts, *val)
	// 	}
	// }

	// ress := []string{}
	// for _, val := range admissionPolicy.Resources {
	// 	if val != nil {
	// 		ress = append(ress, *val)
	// 	}
	// }

	// pbAdmissionsMessage := &admissions.AdmissionMessage{
	// 	Id:         uuid,
	// 	Name:       admissionPolicy.Name,
	// 	Effect:     effect,
	// 	Type:       policyType,
	// 	Principals: princips,
	// 	Actions:    acts,
	// 	Resources:  ress,
	// }

	// region API call
	// response, err := r.apiClient.MakeApiRequest(*createdAdmissionPolicy, "CreateAdmissionPolicy", "POST")
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while trying to POST object: %v", err)
	// }
	// fmt.Printf("Successfully submitted POST request for object %s", *response)
	// endregion API call
	return createdAdmissionPolicy, nil
}

// UpdateAdmissionPolicyActions is the resolver for the updateAdmissionPolicyActions field.
func (r *mutationResolver) UpdateAdmissionPolicyActions(ctx context.Context, id string, admissionPolicyActions []*string) (*model.AdmissionPolicy, error) {
	updateAdmissionPolicyActionsModel := &model.AdmissionPolicyActions{
		ID:      id,
		Actions: admissionPolicyActions,
	}
	fmt.Println(updateAdmissionPolicyActionsModel)
	// encode input struct
	// encodedBlob, err := EncodeToString(updateAdmissionPolicyActionsModel)
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	// }

	// messageData := CommanderMessage{
	// 	Action: "UpdateAdmissionPolicyActions",
	// 	Data: CommanderMessageData{
	// 		MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
	// 		DataBlob:    encodedBlob,
	// 		CommandName: "UpdateAdmissionPolicyActions", // may not need this
	// 	},
	// }
	// updatedAdmissionPolicy := &model.AdmissionPolicy{
	// 	ID:        id,
	// 	Name:      "",
	// 	Effect:    nil,
	// 	Type:      nil,
	// 	Principal: []*string{},
	// 	Actions:   append([]*string{}, admissionPolicyActions...),
	// 	Resources: []*string{},
	// }
	// region API call
	// response, err := r.apiClient.MakeApiRequest(*updatedAdmissionPolicy, "UpdateAdmissionPolicyActions", "PUT")
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while trying to PUT object: %v", err)
	// }
	// fmt.Printf("Successfully submitted PUT request for object %s", *response)
	// region API call
	// return updateAdmissionPolicyActionsModel, nil
	matchingPolicy := model.AdmissionPolicy{}
	for _, v := range r.admissionPolicies {
		if *v.ID == id {
			v.Actions = updateAdmissionPolicyActionsModel.Actions
			matchingPolicy = *v
		}
	}
	return &matchingPolicy, nil
}

// UpdateAdmissionPolicyPrincipals is the resolver for the updateAdmissionPolicyPrincipals field.
func (r *mutationResolver) UpdateAdmissionPolicyPrincipals(ctx context.Context, id string, admissionPolicyPrincipals []*string) (*model.AdmissionPolicy, error) {
	updateAdmissionPolicyPrincipalsModel := &model.AdmissionPolicyPrincipals{
		ID:         id,
		Principals: admissionPolicyPrincipals,
	}
	fmt.Println(updateAdmissionPolicyPrincipalsModel)

	// encode input struct
	// encodedBlob, err := EncodeToString(updateAdmissionPolicyPrincipalsModel)
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	// }

	// messageData := CommanderMessage{
	// 	Action: "UpdateAdmissionPolicyPrincipals",
	// 	Data: CommanderMessageData{
	// 		MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
	// 		DataBlob:    encodedBlob,
	// 		CommandName: "UpdateAdmissionPolicyPrincipals", // may not need this
	// 	},
	// }
	// updatedAdmissionPolicy := &model.AdmissionPolicy{
	// 	ID:        id,
	// 	Name:      "",
	// 	Effect:    nil,
	// 	Type:      nil,
	// 	Principal: append([]*string{}, admissionPolicyPrincipals...),
	// 	Actions:   []*string{},
	// 	Resources: []*string{},
	// }
	// region API call
	// response, err := r.apiClient.MakeApiRequest(*updatedAdmissionPolicy, "UpdateAdmissionPolicyPrincipals", "PUT")
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while trying to PUT object: %v", err)
	// }

	// fmt.Printf("Successfully submitted PUT request for object %s", *response)
	// endregion API call
	// return updateAdmissionPolicyPrincipalsModel, nil
	matchingPolicy := model.AdmissionPolicy{}
	for _, v := range r.admissionPolicies {
		if *v.ID == id {
			v.Principal = updateAdmissionPolicyPrincipalsModel.Principals
			matchingPolicy = *v
		}
	}
	return &matchingPolicy, nil
}

// UpdateAdmissionPolicyResources is the resolver for the updateAdmissionPolicyResources field.
func (r *mutationResolver) UpdateAdmissionPolicyResources(ctx context.Context, id string, admissionPolicyResources []*string) (*model.AdmissionPolicy, error) {
	updateAdmissionPolicyResourcesModel := &model.AdmissionPolicyResources{
		ID:        id,
		Resources: admissionPolicyResources,
	}
	fmt.Println(updateAdmissionPolicyResourcesModel)

	// encode input struct
	// encodedBlob, err := EncodeToString(updateAdmissionPolicyResourcesModel)
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while encoding input object: %v", err)
	// }

	// messageData := CommanderMessage{
	// 	Action: "UpdateAdmissionPolicyResources",
	// 	Data: CommanderMessageData{
	// 		MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
	// 		DataBlob:    encodedBlob,
	// 		CommandName: "UpdateAdmissionPolicyResources", // may not need this
	// 	},
	// }
	// updatedAdmissionPolicy := &model.AdmissionPolicy{
	// 	ID:        id,
	// 	Name:      "",
	// 	Effect:    nil,
	// 	Type:      nil,
	// 	Principal: []*string{},
	// 	Actions:   []*string{},
	// 	Resources: append([]*string{}, admissionPolicyResources...),
	// }
	// region API call
	// response, err := r.apiClient.MakeApiRequest(*updatedAdmissionPolicy, "UpdateAdmissionPolicyResources", "PUT")
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while trying to PUT object: %v", err)
	// }

	// fmt.Printf("Successfully submitted PUT request for object %s", *response)
	// endregion API call
	// return updateAdmissionPolicyResourcesModel, nil
	matchingPolicy := model.AdmissionPolicy{}
	for _, v := range r.admissionPolicies {
		if *v.ID == id {
			v.Resources = updateAdmissionPolicyResourcesModel.Resources
			matchingPolicy = *v
		}
	}
	return &matchingPolicy, nil
}

// DeleteAdmissionPolicy is the resolver for the deleteAdmissionPolicy field.
func (r *mutationResolver) DeleteAdmissionPolicy(ctx context.Context, id string) (*bool, error) {
	// The below logic will be replaced with Commander API call for deletes here, this is temporary for example
	// TODO: Send marshalled JSON object to Commander API for database deletes below here
	var deleted = false
	// messageData := CommanderMessage{
	// 	Action: "DeleteAdmissionPolicy",
	// 	Data: CommanderMessageData{
	// 		MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
	// 		DataBlob:    id,
	// 		CommandName: "DeleteAdmissionPolicy", // may not need this
	// 	},
	// }
	// updatedAdmissionPolicy := &model.AdmissionPolicy{
	// 	ID:        id,
	// 	Name:      "",
	// 	Effect:    nil,
	// 	Type:      nil,
	// 	Principal: []*string{},
	// 	Actions:   []*string{},
	// 	Resources: []*string{},
	// }
	// endregion API call
	// response, err := r.apiClient.MakeApiRequest(*updatedAdmissionPolicy, "DeleteAdmissionPolicy", "DELETE")
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while trying to DELETE object: %v", err)
	// }
	// fmt.Printf("Successfully submitted DELETE request for object %s", *response)
	// endregion API call
	deleted = true

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

// AdmissionPolicy is the resolver for the admissionPolicy field.
func (r *queryResolver) AdmissionPolicy(ctx context.Context, id string) (*model.AdmissionPolicy, error) {
	// panic(fmt.Errorf("not implemented: AdmissionPolicy - admissionPolicy"))
	_, err := uuid.Parse((id))
	if err != nil {
		return nil, fmt.Errorf("invalid identifier: %s", id)
	}
	var tempItem model.AdmissionPolicy
	// for index := range r.admissionPolicies {
	// 	if *r.admissionPolicies[index].ID == id {
	// 		returnItem = *r.admissionPolicies[index]
	// 	}
	// }
	tempItem.ID = &id
	returnItem, err := tempItem.Get()
	if err != nil {
		return nil, err
	}
	return returnItem, nil
}

// AdmissionPolicyRelation is the resolver for the admissionPolicyRelation field.
func (r *queryResolver) AdmissionPolicyRelation(ctx context.Context, principal *string, action *string, resourceID *string) (*model.AdmissionPolicyRelation, error) {
	tempItem := &model.AdmissionPolicyRelation{
		Principal:  *principal,
		Action:     action,
		ResourceID: resourceID,
	}
	returnItem, err := tempItem.GetByPrincipalActionResource()
	if err != nil {
		return nil, err
	}
	return returnItem, nil
}

// AdmissionPolicyRelations is the resolver for the admissionPolicyRelations field.
func (r *queryResolver) AdmissionPolicyRelations(ctx context.Context, principal *string, action *string, resourceID *string) ([]*model.AdmissionPolicyRelation, error) {
	panic(fmt.Errorf("not implemented: AdmissionPolicyRelations - admissionPolicyRelations"))
}

// AdmissionPolicyAuthorizationCheck is the resolver for the admissionPolicyAuthorizationCheck field.
func (r *queryResolver) AdmissionPolicyAuthorizationCheck(ctx context.Context, principal string, action string, resourceID string, ttl *string) (*model.AdmissionPolicyAuthorization, error) {
	panic(fmt.Errorf("not implemented: AdmissionPolicyAuthorizationCheck - admissionPolicyAuthorizationCheck"))
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
