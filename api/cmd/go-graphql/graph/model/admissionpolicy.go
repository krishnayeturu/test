package model

import (
	"fmt"

	database "gitlab.com/2ndwatch/microservices/ms-admissions-service/api/pkg/services"
)

func (admissionPolicy AdmissionPolicy) Insert() (int64, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(`INSERT INTO AdmissionPolicy (UUID, Name, Type) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(admissionPolicy.ID, admissionPolicy.Name, admissionPolicy.Type)
	if err != nil {
		return -1, err
	} else {
		insert_id, err := result.LastInsertId()
		if err != nil {
			return -1, err
		}
		fmt.Printf("Successfully inserted supplied admission policy %d", insert_id)

		// Test below thoroughly as it seems clunky
		for _, principal := range admissionPolicy.Principal {
			if principal != nil {
				for _, resource := range admissionPolicy.Resources {
					if resource != nil {
						for _, action := range admissionPolicy.Actions {
							if action != nil {
								createAdmissionPolicyStatement := &AdmissionPolicyStatement{
									PolicyID:   *admissionPolicy.ID,
									Effect:     *admissionPolicy.Effect,
									Principal:  *principal,
									Action:     action,
									ResourceID: resource,
								}
								_, err := createAdmissionPolicyStatement.Insert()
								if err != nil {
									return -1, err
								}
							}
						}
					}
				}
			}
		}
		return insert_id, nil
	}
}

func (admissionPolicy AdmissionPolicy) UpdatePolicyStatements() (*AdmissionPolicy, error) {
	database.ConnectDB()
	// first clear out old policystatements
	statement, err := database.Db.Prepare(fmt.Sprintf("delete from AdmissionPolicyStatement where PolicyUuid = '%s'", *admissionPolicy.ID))
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	_, err = statement.Exec()
	if err != nil {
		return nil, err
	}
	// now add new policystatements
	for _, principal := range admissionPolicy.Principal {
		if principal != nil {
			for _, resource := range admissionPolicy.Resources {
				if resource != nil {
					for _, action := range admissionPolicy.Actions {
						if action != nil {
							createAdmissionPolicyStatement := &AdmissionPolicyStatement{
								PolicyID:   *admissionPolicy.ID,
								Effect:     *admissionPolicy.Effect,
								Principal:  *principal,
								Action:     action,
								ResourceID: resource,
							}
							_, err := createAdmissionPolicyStatement.Insert()
							if err != nil {
								return nil, err
							}
						}
					}
				}
			}
		}
	}
	result, err := admissionPolicy.Get()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (admissionPolicy AdmissionPolicy) Get() (*AdmissionPolicy, error) {
	// dbStatement := fmt.Sprintf("select ap.UUID as Id, ap.Name, ap.Type as AdmissionPolicyType, apr.Principal as Principal, apr.Action as Action, apr.ResourceId as Resource from AdmissionPolicy ap join AdmissionPolicyStatement apr on ap.Id = apr.PolicyId where ap.UUID = '%s'", *admissionPolicy.ID)
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("select ap.UUID as ID, ap.Name, ap.Type from AdmissionPolicy ap where ap.UUID = '%s'", *admissionPolicy.ID))
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	var result AdmissionPolicy
	for rows.Next() {
		var res AdmissionPolicy
		err := rows.Scan(&res.ID, &res.Name, &res.Type)
		// TODO: this will require ORM addition or join statements to retrieve principal, actions, and resources
		if err != nil {
			return nil, err
		}
		result = res
		admissionPolicyStatements, err := GetAdmissionPolicyStatementsForPolicyUuid(*admissionPolicy.ID)
		if err != nil {
			return nil, err
		}
		principals := []*string{}
		actions := []*string{}
		resources := []*string{}
		for index := range admissionPolicyStatements {
			principals = append(principals, &admissionPolicyStatements[index].Principal)
			actions = append(actions, admissionPolicyStatements[index].Action)
			resources = append(resources, admissionPolicyStatements[index].ResourceID)
		}
		if len(admissionPolicyStatements) > 0 {
			result.Effect = &admissionPolicyStatements[len(admissionPolicyStatements)-1].Effect
		} else {
			result.Effect = nil
		}
		result.Principal = principals
		result.Actions = actions
		result.Resources = resources
	}
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (admissionPolicy AdmissionPolicy) Delete() (bool, error) {
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("delete from AdmissionPolicy ap where ap.UUID = '%s'", *admissionPolicy.ID))
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
