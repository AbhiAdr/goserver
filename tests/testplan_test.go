package main

import (
	"fmt"
	"go_mjolnir/models"
	// "go_mjolnir/structs"
	//"reflect"
	"testing"
)

// func TestTestplan_create_update_get(t *testing.T) {

// 	models.InitDB("devopsdb:6bCDn.nx@tcp(gomjolnir-dev.cepqlfnboyp8.us-west-2.rds.amazonaws.com:3306)/mjolnir_dev?charset=utf8")

// 	// Create New Testplan
// 	uuid := models.Pseudo_uuid()
// 	testplan := structs.Testplan{Name: "Name_" + uuid, Desc: "Desc_" + uuid, Modified_By: "1", Team_id: "TEAM-UUID-01"}
// 	testplan, err := models.Testplan_create(testplan)
// 	fmt.Printf("1.Testplan Create : %v\n", testplan)

// 	if err != nil {
// 		fmt.Println("Error message", err)
// 		fmt.Printf("Failed Update : %v\n", testplan)

// 	}

// 	// Update Testplan created previously

// 	scenario := `[{
// 						"TC01": { "start": 10, "delay": 0, "rampup": 30, "run": 3600, "rampdown": 0},
// 						"TC02": { "start": 10, "delay": 0, "rampup": 30, "run": 3600, "rampdown": 0}
// 					}]`

// 	testplan.Jprofile_id = "PROFILE-UUID-01"
// 	testplan.Scenario = scenario
// 	testplan, err = models.Testplan_update(testplan)
// 	fmt.Printf("2.Testplan Update : %v\n", testplan)

// 	if err != nil {
// 		fmt.Println("Error message", err)
// 	}

// 	// Get Testplan.
// 	testplan, err = models.Testplan(testplan)
// 	fmt.Printf("3. Get Testplan : %v\n", testplan)

// }

func TestTestplans(t *testing.T) {
	models.InitDB("devopsdb:6bCDn.nx@tcp(gomjolnir-dev.cepqlfnboyp8.us-west-2.rds.amazonaws.com:3306)/mjolnir_dev?charset=utf8")

	// List of Testplans
	data, err := models.Testplans("USER-UUID-01")
	fmt.Printf("1.Testplan List : %v\n", data)

	if err != nil {
		fmt.Println("Error message", err)
		fmt.Printf("Failed Update : %v\n", data)

	}

}
