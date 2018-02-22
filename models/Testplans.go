package models

import (
	"errors"
	"fmt"
	"go_mjolnir/structs"
	// "reflect"
)

//func Testplan_create(name string, desc string, modified_by string, team_id string) structs.Testplanbasic {
func Testplan_create(data structs.Testplan) (structs.Testplan, error) {

	data.Id = Pseudo_uuid()

	stmt, err := db.Prepare("INSERT tbl_testplan(id,name,description,last_modified_by,team_id) VALUES(?,?,?,?,?)")

	checkErr(err)
	_, err = stmt.Exec(data.Id, data.Name, data.Desc, data.Modified_By, data.Team_id)

	checkErr(err)

	if err != nil {
		return data, errors.New("Unable to create user")
	}
	return data, nil
}

func Testplans(userId string) (structs.Testplans, error) {

	var data structs.Testplans
	var testplan structs.Testplan
	// Get all test plans belongs to the teams in which user is
	query := `
	SELECT t1.id, t1.name AS name , t2.name AS team , t1.lg_count , t3.name AS profile , t4.firstname as user FROM tbl_testplan t1 , tbl_team t2 , tbl_jmeter_profile t3 , tbl_users t4 WHERE t1.team_id in ( SELECT team_id FROM tbl_users_team WHERE users_id IN ( SELECT id FROM tbl_users WHERE id = ?) ) AND t2.id = t1.team_id AND t1.jmeter_profile_id = t3.id;
	`
	rows, err := db.Query(query, userId)

	checkErr(err)

	for rows.Next() {

		err = rows.Scan(&testplan.Id, &testplan.Name, &testplan.Team_id, &testplan.Lg_count, &testplan.Jprofile_id, &testplan.Modified_By)
		checkErr(err)
		data = append(data, testplan)
	}
	fmt.Println(data)
	if err != nil {
		return data, errors.New("Unable to create Testplan")
	}
	return data, nil
}

func Testplan(data structs.Testplan) (structs.Testplan, error) {

	stmt, err := db.Prepare("SELECT t1.id, t1.name,  t1.description , t1.lg_count, t1.scenario , t1.last_modified_on , t1.last_modified_by, t2.name  FROM tbl_testplan t1 , tbl_jmeter_profile t2 where t1.id = ? and t1.jmeter_profile_id = t2.id")

	checkErr(err)
	defer stmt.Close()

	rows, err := stmt.Query(data.Id)
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&data.Id, &data.Name, &data.Desc, &data.Lg_count, &data.Scenario, &data.Modified_On, &data.Modified_By, &data.Jprofile_id)
		checkErr(err)
	}

	if err != nil {
		return data, errors.New("Unable to Fetch Testplan")
	}
	return data, nil
}

func Testplan_update(data structs.Testplan) (structs.Testplan, error) {

	stmt, err := db.Prepare("UPDATE tbl_testplan set name = ?  , jmeter_profile_id = ? , last_modified_by = ?  , scenario = ? where id = ? ")

	if err != nil {
		panic(err.Error())
	}

	checkErr(err)

	_, err = stmt.Exec(data.Name, data.Jprofile_id, data.Modified_By, data.Scenario, data.Id)

	checkErr(err)

	if err != nil {
		return data, errors.New("Unable to Fetch Testplan")
	}
	return data, nil

}
