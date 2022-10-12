package model

import (
	"fmt"

	database "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/services"
)

func (admissionPolicyStatement AdmissionPolicyStatement) Insert() (int64, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(`INSERT INTO AdmissionPolicyStatement (PolicyUuid, Effect, Principal, Action, Resourceid) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(admissionPolicyStatement.PolicyID, admissionPolicyStatement.Effect, admissionPolicyStatement.Principal, admissionPolicyStatement.Action, admissionPolicyStatement.ResourceID)
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

func (admissionPolicyStatement AdmissionPolicyStatement) Get() (*AdmissionPolicyStatement, error) {
	// dbStatement := fmt.Sprintf("select ap.UUID as Id, ap.Name, ap.Type as AdmissionPolicyType, apr.Principal as Principal, apr.Action as Action, apr.ResourceId as Resource from AdmissionPolicy ap join AdmissionPolicyStatement apr on ap.Id = apr.PolicyId where ap.UUID = '%s'", *admissionPolicy.ID)
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("select Id, PolicyUuid, Effect, Principal, Action, ResourceId from AdmissionPolicyStatement where Id = '%s'", *admissionPolicyStatement.ID))
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	var result AdmissionPolicyStatement
	for rows.Next() {
		var res AdmissionPolicyStatement
		err := rows.Scan(&res.ID, &res.PolicyID, &res.Effect, &res.Principal, &res.Action, &res.ResourceID)
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

func GetAdmissionPolicyStatementsForPolicyUuid(policyUuid string) ([]AdmissionPolicyStatement, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("select Id, PolicyUuid, Effect, Principal, Action, ResourceId from AdmissionPolicyStatement where PolicyUuid = '%s'", policyUuid))
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	var result []AdmissionPolicyStatement
	for rows.Next() {
		var res AdmissionPolicyStatement
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

func (admissionPolicyStatement AdmissionPolicyStatement) GetByPrincipalActionResource() (*AdmissionPolicyStatement, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("select Id, PolicyUuid, Effect, Principal, Action, ResourceId from AdmissionPolicyStatement where Principal = '%s' and Action = '%s' and ResourceId = '%s'", admissionPolicyStatement.Principal, *admissionPolicyStatement.Action, *admissionPolicyStatement.ResourceID))
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	var result AdmissionPolicyStatement
	for rows.Next() {
		var res AdmissionPolicyStatement
		err := rows.Scan(&res.ID, &res.PolicyID, &res.Effect, &res.Principal, &res.Action, &res.ResourceID)
		if err != nil {
			return nil, err
		}
		result = res
	}
	if err != nil {
		return nil, err
	}
	if result.ID == nil {
		return nil, nil
	}
	return &result, nil
}
