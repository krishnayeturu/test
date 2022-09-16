package graph

//go:generate go run github.com/99designs/gqlgen generate
import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"gitlab.com/2ndwatch/microservices/ms-admissions-service/ms-admissions-api/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DbClient                   *dynamodb.Client
	todos                      []*model.Todo
	admissionPolicies          []*model.AdmissionPolicy
	principalAdmissionPolicies []*model.PrincipalAdmissionPolicy
	AdmissionPolicy            *model.AdmissionPolicy
	PrincipalAdmissionPolicy   *model.PrincipalAdmissionPolicy
}
