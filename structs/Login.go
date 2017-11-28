package structs

type Login struct {
	Status   string      `json:"status"`
	Data     UserDetails `json:"data"`
	Login_id int64       `json:"id"`
}

type Login_err struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
