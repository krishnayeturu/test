package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/ms-admissions-api/graph/generated"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/ms-admissions-api/graph/model"
)

// CreateAdmissionPolicy is the resolver for the createAdmissionPolicy field.
func (r *mutationResolver) CreateAdmissionPolicy(ctx context.Context, admissionPolicy model.AdmissionPolicyInput) (*model.AdmissionPolicy, error) {
	uuid := strings.Replace(uuid.New().String(), "-", "", -1)

	createdAdmissionPolicy := &model.AdmissionPolicy{
		ID:        uuid,
		Name:      admissionPolicy.Name,
		Effect:    admissionPolicy.Effect,
		Type:      admissionPolicy.Type,
		Principal: append([]*string{}, admissionPolicy.Principal...),
		Actions:   append([]*string{}, admissionPolicy.Actions...),
		Resources: append([]*string{}, admissionPolicy.Resources...),
	}
	// Do database inserts here
	r.admissionPolicies = append(r.admissionPolicies, createdAdmissionPolicy)
	return createdAdmissionPolicy, nil
}

// UpdateAdmissionPolicy is the resolver for the updateAdmissionPolicy field.
func (r *mutationResolver) UpdateAdmissionPolicy(ctx context.Context, id string, admissionPolicy model.AdmissionPolicyInput) (*model.AdmissionPolicy, error) {
	_, err := uuid.Parse(id)

	if err != nil {
		return nil, fmt.Errorf("invalid admission policy identifier: %s", id)
	}
	// _, err := r.AdmissionPolicy.Update(id, admissionPolicy.)
	updatedAdmissionPolicy := &model.AdmissionPolicy{
		Name:      admissionPolicy.Name,
		Effect:    admissionPolicy.Effect,
		Type:      admissionPolicy.Type,
		Principal: append([]*string{}, admissionPolicy.Principal...),
		Actions:   append([]*string{}, admissionPolicy.Actions...),
		Resources: append([]*string{}, admissionPolicy.Resources...),
	}
	// Make Commander API call for updates here, this is temporary for example
	r.admissionPolicies = append(r.admissionPolicies, updatedAdmissionPolicy)
	return updatedAdmissionPolicy, nil
}

// DeleteAdmissionPolicy is the resolver for the deleteAdmissionPolicy field.
func (r *mutationResolver) DeleteAdmissionPolicy(ctx context.Context, admissionPolicy model.AdmissionPolicyInput) (*bool, error) {
	// The below logic will be replaced with Commander API call for deletes here, this is temporary for example
	var admissionPolicies = []*model.AdmissionPolicy{}
	var deleted = false
	for i := 0; i < len(r.admissionPolicies); i++ {
		if r.admissionPolicies[i].ID != *admissionPolicy.ID {
			admissionPolicies = append(admissionPolicies, r.admissionPolicies[i])
		} else {
			deleted = true
		}
	}
	r.admissionPolicies = admissionPolicies
	return &deleted, nil
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	uuid := strings.Replace(uuid.New().String(), "-", "", -1)

	todo := &model.Todo{
		Text: input.Text,
		ID:   uuid,
		User: &model.User{ID: input.UserID, Name: fmt.Sprintf("user '%s'", input.UserID)},
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

// PrincipalAdmissionPolicies is the resolver for the principalAdmissionPolicies field.
func (r *queryResolver) PrincipalAdmissionPolicies(ctx context.Context) ([]*model.PrincipalAdmissionPolicy, error) {
	return r.principalAdmissionPolicies, nil
}

// User is the resolver for the user field.
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.User.ID, Name: obj.User.Name}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
