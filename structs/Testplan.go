package structs

type Testplan struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Lg_count    string `json:"lg_count"`
	Modified_On string `json:"modified_On"`
	Modified_By string `json:"modified_By"`
	Team_id     string `json:"team_id"`
	Jprofile_id string `json:"jprofile_id"`
	Scenario    string `json:"Scenario"`
	Runs        string `json:"runs"`
}

type Testplans []Testplan
