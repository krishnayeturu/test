package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
	"time"

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

	// region Database Operations
	insert_id, err := createdAdmissionPolicy.Insert()
	if err != nil {
		return nil, err
	}

	fmt.Println(insert_id)
	// enddregion Database Operations

	// region API call
	response, err := r.apiClient.MakeApiRequest(*createdAdmissionPolicy, "CreateAdmissionPolicy", "POST")
	if err != nil {
		return nil, fmt.Errorf("encountered an error while trying to POST object: %v", err)
	}
	fmt.Printf("Successfully submitted POST request for object %s", *response)
	// endregion API call
	return createdAdmissionPolicy, nil
}

// UpdateAdmissionPolicy is the resolver for the updateAdmissionPolicy field.
func (r *mutationResolver) UpdateAdmissionPolicy(ctx context.Context, admissionPolicyUpdates model.AdmissionPolicyUpdateInput) (*model.AdmissionPolicy, error) {
	updatedAdmissionPolicy := &model.AdmissionPolicy{
		ID:        &admissionPolicyUpdates.ID,
		Effect:    &admissionPolicyUpdates.Effect,
		Principal: append([]*string{}, admissionPolicyUpdates.Principal...),
		Actions:   append([]*string{}, admissionPolicyUpdates.Actions...),
		Resources: append([]*string{}, admissionPolicyUpdates.Resources...),
	}
	result, err := updatedAdmissionPolicy.UpdatePolicyStatements()
	if err != nil {
		return nil, err
	}

	// region API call
	// response, err := r.apiClient.MakeApiRequest(*updatedAdmissionPolicy, "UpdateAdmissionPolicy", "PUT")
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while trying to PUT object: %v", err)
	// }
	// fmt.Printf("Successfully submitted PUT request for object %s", *response)
	// endregion API call

	return result, nil
}

// DeleteAdmissionPolicy is the resolver for the deleteAdmissionPolicy field.
func (r *mutationResolver) DeleteAdmissionPolicy(ctx context.Context, id string) (*bool, error) {
	// The below logic will be replaced with Commander API call for deletes here, this is temporary for example
	// TODO: Send marshalled JSON object to Commander API for database deletes below here
	var deleted = false
	deletedAdmissionPolicy := &model.AdmissionPolicy{
		ID: &id,
	}
	_, err := deletedAdmissionPolicy.Delete()
	if err != nil { // NOOP - nothing to delete
		return &deleted, nil
	}

	// region API call
	// response, err := r.apiClient.MakeApiRequest(*deletedAdmissionPolicy, "DeleteAdmissionPolicy", "DELETE")
	// if err != nil {
	// 	return nil, fmt.Errorf("encountered an error while trying to DELETE object: %v", err)
	// }
	// fmt.Printf("Successfully submitted DELETE request for object %s", *response)
	// endregion API call
	deleted = true

	return &deleted, nil
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
	tempItem.ID = &id
	returnItem, err := tempItem.Get()
	if err != nil {
		return nil, err
	}
	return returnItem, nil
}

// AdmissionPolicyStatement is the resolver for the admissionPolicyStatement field.
func (r *queryResolver) AdmissionPolicyStatement(ctx context.Context, principal *string, action *string, resourceID *string) (*model.AdmissionPolicyStatement, error) {
	tempItem := &model.AdmissionPolicyStatement{
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

// AdmissionPolicyStatements is the resolver for the admissionPolicyStatements field.
func (r *queryResolver) AdmissionPolicyStatements(ctx context.Context, principal *string, action *string, resourceID *string) ([]*model.AdmissionPolicyStatement, error) {
	panic(fmt.Errorf("not implemented: AdmissionPolicyStatements - admissionPolicyStatements"))
}

// AdmissionPolicyAuthorizationCheck is the resolver for the admissionPolicyAuthorizationCheck field.
func (r *queryResolver) AdmissionPolicyAuthorizationCheck(ctx context.Context, principal string, action string, resourceID string, ttl string) (*model.AdmissionPolicyAuthorization, error) {
	tempItem := &model.AdmissionPolicyStatement{
		Principal:  principal,
		Action:     &action,
		ResourceID: &resourceID,
	}
	fetchedItem, err := tempItem.GetByPrincipalActionResource()
	if err != nil {
		return nil, err
	}
	if ttl == "" {
		// # default
		ttl = "1h0m0s"
	}
	parsedTtl, err := time.ParseDuration(ttl)
	if err != nil {
		return nil, fmt.Errorf("TTL input value must be in hour minute second format, example: <A>h<B>m<C>s")
	}
	ttl = parsedTtl.String()
	admissionPolicyAuthorized := &model.AdmissionPolicyAuthorization{
		Principal:           principal,
		AuthorizationResult: fetchedItem != nil,
		ExpireTime:          &ttl,
	}
	return admissionPolicyAuthorized, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
