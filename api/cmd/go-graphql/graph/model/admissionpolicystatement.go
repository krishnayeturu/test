package model

import (
	"fmt"

	database "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/services"
)

func (admissionPolicyStatement AdmissionPolicyStatement) Insert() (int64, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(`insert ignore into AdmissionPolicyStatement (PolicyUuid, Effect, Principal, Action, Resourceid) values (?, ?, ?, ?, ?)`)
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

func (admissionPolicyStatement AdmissionPolicyStatement) Get() (*AdmissionPolicyStatement, error) {
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

func (admissionPolicyStatement AdmissionPolicyStatement) Delete() (bool, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("delete from AdmissionPolicyStatement aps where aps.Id = '%s'", *admissionPolicyStatement.ID))
	if err != nil {
		return false, err
	}
	defer statement.Close()
	_, err = statement.Query()
	if err != nil {
		return false, err
	}
	return true, nil
}
