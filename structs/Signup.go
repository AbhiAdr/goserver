package structs

type Signup struct {
	Status   string      `json:"status"`
	Data     UserDetails `json:"data"`
	Login_id int64       `json:"id"`
}

type Signup_err struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
