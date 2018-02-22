package structs

type Login struct {
	Status   string `json:"status"`
	Data     User   `json:"data"`
	Login_id string `json:"id"`
}

type Login_err struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
