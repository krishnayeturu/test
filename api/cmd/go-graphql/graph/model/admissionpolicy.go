package model

import (
	"fmt"
	"strconv"

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
								createAdmissionPolicyRelation := &AdmissionPolicyRelation{
									PolicyID:   strconv.FormatInt(insert_id, 10),
									Effect:     *admissionPolicy.Effect,
									Principal:  *principal,
									Action:     action,
									ResourceID: resource,
								}
								_, err := createAdmissionPolicyRelation.Insert()
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

func (admissionPolicy AdmissionPolicy) Update() (int64, error) {
	// TODO: below logic is placeholder, need to determine better approach for upsert logic
	// database.ConnectDB()
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
								createAdmissionPolicyRelation := &AdmissionPolicyRelation{
									PolicyID:   strconv.FormatInt(insert_id, 10),
									Effect:     *admissionPolicy.Effect,
									Principal:  *principal,
									Action:     action,
									ResourceID: resource,
								}
								_, err := createAdmissionPolicyRelation.Insert()
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

func (admissionPolicy AdmissionPolicy) Get() (*AdmissionPolicy, error) {
	// dbStatement := fmt.Sprintf("select ap.UUID as Id, ap.Name, ap.Type as AdmissionPolicyType, apr.Principal as Principal, apr.Action as Action, apr.ResourceId as Resource from AdmissionPolicy ap join AdmissionPolicyRelation apr on ap.Id = apr.PolicyId where ap.UUID = '%s'", *admissionPolicy.ID)
	database.ConnectDB()
	statement, err := database.Db.Prepare(fmt.Sprintf("select ap.UUID as ID, ap.Name, ap.Type, apr.Effect as Effect, apr.Principal as Principal from AdmissionPolicy ap join AdmissionPolicyRelation apr on ap.Id = apr.PolicyId where UUID = '%s'", *admissionPolicy.ID))
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
		err := rows.Scan(&res.ID, &res.Name, &res.Type, &res.Effect)
		// TODO: this will require ORM addition or join statements to retrieve principal, actions, and resources
		if err != nil {
			return nil, err
		}
		result = res
	}
	// var fetchedAdmissionPolicy AdmissionPolicy
	// result, err := database.Db.Query(dbStatement)
	// defer result.Close()
	// https://zetcode.com/golang/mysql/
	// err := database.Db.QueryRow(dbStatement).Scan(&fetchedAdmissionPolicy)
	if err != nil {
		return nil, err
	}
	return &result, nil

}
