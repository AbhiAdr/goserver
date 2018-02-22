package main

import (
	"fmt"
	"goserver/models"
	"goserver/structs"
	//"reflect"
	"testing"
)

func TestUser_login_logout(t *testing.T) {

	models.InitDB("devopsdb:6bCDn.nx@tcp(gomjolnir-dev.cepqlfnboyp8.us-west-2.rds.amazonaws.com:3306)/mjolnir_dev?charset=utf8")

	// Login User
	//uuid := models.Pseudo_uuid()
	user := structs.User{Email: "admin@abc.com"}

	data, err := models.User_get_by_username(user)
	fmt.Printf("1.User Login : %v\n", data)
	if err != nil {
		fmt.Println("Error message", err)
		fmt.Printf("Failed Update : %v\n", user)

	}

	data2, err := models.Users()
	fmt.Printf("1.User list : %v\n", data2)

	/*
	    	// Update Testplan created previously

	    	scenario := `[{
	    						"TC01": { "start": 10, "delay": 0, "rampup": 30, "run": 3600, "rampdown": 0},
	    						"TC02": { "start": 10, "delay": 0, "rampup": 30, "run": 3600, "rampdown": 0}
	    					}]`

	    	testplan.Jprofile_id = "PROFILE-UUID-01"
	   // 	testplan.Scenario = scenario
	   // 	testplan, err = models.Testplan_update(testplan)
	   // 	fmt.Printf("2.Testplan Update : %v\n", testplan)

	   // 	if err != nil {
	   // 		fmt.Println("Error message", err)
	   // 	}

	   // 	// Get Testplan.
	   // 	testplan, err = models.Testplan(testplan)
	   // 	fmt.Printf("3. Get Testplan : %v\n", testplan)
	*/
}

/*
func TestUsers(t *testing.T) {
	models.InitDB("devopsdb:6bCDn.nx@tcp(gomjolnir-dev.cepqlfnboyp8.us-west-2.rds.amazonaws.com:3306)/mjolnir_dev?charset=utf8")

	// List of Testplans
	data, err := models.Users("1")
	fmt.Printf("1.Testplan Create : %v\n", data)

	if err != nil {
		fmt.Println("Error message", err)
		fmt.Printf("Failed Update : %v\n", data)

	}

}
*/
