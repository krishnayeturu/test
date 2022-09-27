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
		ID:        policyUuid,
		Name:      admissionPolicy.Name,
		Effect:    admissionPolicy.Effect,
		Type:      admissionPolicy.Type,
		Principal: append([]*string{}, admissionPolicy.Principal...),
		Actions:   append([]*string{}, admissionPolicy.Actions...),
		Resources: append([]*string{}, admissionPolicy.Resources...),
	}
	// TODO: Send marshalled JSON object to Commander API for database inserts here
	r.admissionPolicies = append(r.admissionPolicies, createdAdmissionPolicy)
	// messageData := &commanderclient.CommanderMessage{
	// 	Action: "CreateAdmissionPolicy",
	// 	Data: commanderclient.CommanderMessageData{
	// 		MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
	// 		DataBlob:    *createdAdmissionPolicy,
	// 		CommandName: "CreateAdmissionPolicy",
	// 	},
	// }
	// _, err = commanderclient.MakeApiRequest(messageData, "POST")
	// if err != nil {
	// 	return fmt.Errorf("encountered an error while trying to POST object: %v", err)
	// }

	// fmt.Printf("Successfully submitted POST request for object %s", data)
	return createdAdmissionPolicy, nil
}

// DeleteAdmissionPolicy is the resolver for the deleteAdmissionPolicy field.
func (r *mutationResolver) DeleteAdmissionPolicy(ctx context.Context, id string) (*bool, error) {
	// The below logic will be replaced with Commander API call for deletes here, this is temporary for example
	// TODO: Send marshalled JSON object to Commander API for database deletes below here
	var deleted = false
	// messageData := &commanderclient.CommanderMessage{
	// 	Action: "DeleteAdmissionPolicy",
	// 	Data: CommanderMessageData{
	// 		MessageUUID: strings.Replace(uuid.New().String(), "-", "", -1),
	// 		DataBlob:    *id,
	// 		CommandName: "DeleteAdmissionPolicy", // may not need this
	// 	},
	// }
	// _, err = commanderclient.MakeApiRequest(messageData, "DELETE")
	// if err != nil {
	// 	return fmt.Errorf("encountered an error while trying to DELETE object: %v", err)
	// }
	// deleted = true

	// fmt.Printf("Successfully submitted DELETE request for object %s", data)
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
