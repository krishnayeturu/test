package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/api/cmd/go-graphql/graph/model"
)

type Resolver struct {
	todos             []*model.Todo
	admissionPolicies []*model.AdmissionPolicy
	AdmissionPolicy   *model.AdmissionPolicy
	apiClient         *CommanderClient
}
