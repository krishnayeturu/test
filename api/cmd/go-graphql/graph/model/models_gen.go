// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AdmissionPolicy struct {
	ID        *string              `json:"id"`
	Name      string               `json:"name"`
	Effect    *Effect              `json:"effect"`
	Type      *AdmissionPolicyType `json:"type"`
	Principal []*string            `json:"principal"`
	Actions   []*string            `json:"actions"`
	Resources []*string            `json:"resources"`
}

type AdmissionPolicyAuthorization struct {
	Principal           string  `json:"principal"`
	AuthorizationResult bool    `json:"authorizationResult"`
	ExpireTime          *string `json:"expireTime"`
}

type AdmissionPolicyInput struct {
	ID        *string              `json:"id"`
	Name      string               `json:"name"`
	Effect    *Effect              `json:"effect"`
	Type      *AdmissionPolicyType `json:"type"`
	Principal []*string            `json:"principal"`
	Actions   []*string            `json:"actions"`
	Resources []*string            `json:"resources"`
}

type AdmissionPolicyRelation struct {
	ID         *string `json:"id"`
	PolicyID   string  `json:"policyId"`
	Effect     Effect  `json:"effect"`
	Principal  string  `json:"principal"`
	Action     *string `json:"action"`
	ResourceID *string `json:"resourceId"`
}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AdmissionPolicyType string

const (
	AdmissionPolicyTypeCredential AdmissionPolicyType = "CREDENTIAL"
)

var AllAdmissionPolicyType = []AdmissionPolicyType{
	AdmissionPolicyTypeCredential,
}

func (e AdmissionPolicyType) IsValid() bool {
	switch e {
	case AdmissionPolicyTypeCredential:
		return true
	}
	return false
}

func (e AdmissionPolicyType) String() string {
	return string(e)
}

func (e *AdmissionPolicyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AdmissionPolicyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AdmissionPolicyType", str)
	}
	return nil
}

func (e AdmissionPolicyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Effect string

const (
	EffectAllow Effect = "ALLOW"
	EffectDeny  Effect = "DENY"
)

var AllEffect = []Effect{
	EffectAllow,
	EffectDeny,
}

func (e Effect) IsValid() bool {
	switch e {
	case EffectAllow, EffectDeny:
		return true
	}
	return false
}

func (e Effect) String() string {
	return string(e)
}

func (e *Effect) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Effect(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Effect", str)
	}
	return nil
}

func (e Effect) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
