package model

import (
	"fmt"

	database "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/services"
)

func (admissionPolicyRelation AdmissionPolicyRelation) Insert() (int64, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(`INSERT INTO AdmissionPolicyRelation (PolicyId, Effect, Principal, Action, Resourceid) VALUES (?, ?, ?, ?, ?)`)
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
	statement, err := database.Db.Prepare(fmt.Sprintf("select ap.Id as ID, ap.PolicyId, ap.Effect, apr.Principal as Principal, apr.Action, apr.ResourceId from AdmissionPolicyRelation where Id = '%s'", admissionPolicyRelation.ID))
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
	// var fetchedAdmissionPolicy AdmissionPolicy
	// result, err := database.Db.Query(dbStatement)
	// defer result.Close()
	// err := database.Db.QueryRow(dbStatement).Scan(&fetchedAdmissionPolicy)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
