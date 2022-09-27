package model

type UpdateAdmissionPolicyActions struct {
	ID      string    `json:"id"`
	Actions []*string `json:"actions"`
}

type UpdateAdmissionPolicyPrincipals struct {
	ID         string    `json:"id"`
	Principals []*string `json:"principals"`
}

type UpdateAdmissionPolicyResources struct {
	ID        string    `json:"id"`
	Resources []*string `json:"resources"`
}
