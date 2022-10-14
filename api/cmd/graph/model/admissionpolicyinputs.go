package model

type AdmissionPolicyActions struct {
	ID      string    `json:"id"`
	Actions []*string `json:"actions"`
}

type AdmissionPolicyPrincipals struct {
	ID         string    `json:"id"`
	Principals []*string `json:"principals"`
}

type AdmissionPolicyResources struct {
	ID        string    `json:"id"`
	Resources []*string `json:"resources"`
}
