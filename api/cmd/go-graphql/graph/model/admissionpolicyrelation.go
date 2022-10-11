package model

import (
	"fmt"

	database "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/services"
)

func (admissionPolicyRelation AdmissionPolicyRelation) Insert() (int64, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(`INSERT INTO AdmissionPolicyRelation (PolicyUuid, Effect, Principal, Action, Resourceid) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(admissionPolicyRelation.PolicyID, admissionPolicyRelation.Effect, admissionPolicyRelation.Principal, admissionPolicyRelation.Action, admissionPolicyRelation.ResourceID)
	if err != nil {
		return -1, err
	} else {
		insert_id, err := result.LastInsertId()
		if err != nil {
			return -1, err
		}
		fmt.Printf("Successfully inserted supplied admission policy relation %d", insert_id)
		return insert_id, nil
	}
}

// func (admissionPolicy AdmissionPolicy) Update() (int64, error) {
// 	// TODO: below logic is placeholder, need to determine better approach for upsert logic
// 	// database.ConnectDB()
// }

func (admissionPolicyRelation AdmissionPolicyRelation) Get() (*AdmissionPolicyRelation, error) {
	// dbStatement := fmt.Sprintf("select ap.UUID as Id, ap.Name, ap.Type as AdmissionPolicyType, apr.Principal as Principal, apr.Action as Action, apr.ResourceId as Resource from AdmissionPolicy ap join AdmissionPolicyRelation apr on ap.Id = apr.PolicyId where ap.UUID = '%s'", *admissionPolicy.ID)
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("select Id, PolicyUuid, Effect, Principal, Action, ResourceId from AdmissionPolicyRelation where Id = '%s'", *admissionPolicyRelation.ID))
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	var result AdmissionPolicyRelation
	for rows.Next() {
		var res AdmissionPolicyRelation
		err := rows.Scan(&res.ID, &res.PolicyID, &res.Effect, &res.Principal, &res.Action, &res.ResourceID)
		// TODO: this will require ORM addition or join statements to retrieve principal, actions, and resources
		if err != nil {
			return nil, err
		}
		result = res
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func GetAdmissionPolicyRelationsForPolicyUuid(policyUuid string) ([]AdmissionPolicyRelation, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("select Id, PolicyUuid, Effect, Principal, Action, ResourceId from AdmissionPolicyRelation where PolicyUuid = '%s'", policyUuid))
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	var result []AdmissionPolicyRelation
	for rows.Next() {
		var res AdmissionPolicyRelation
		err := rows.Scan(&res.ID, &res.PolicyID, &res.Effect, &res.Principal, &res.Action, &res.ResourceID)
		if err != nil {
			return nil, err
		}
		result = append(result, res)
	}
	if err != nil {
		return nil, err
	}
	return result, nil
}
