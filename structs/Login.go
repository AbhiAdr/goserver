package structs

type UserDetails struct {
	Id     int64  `json:"id"`
	Fn     string `json:"fn"`
	Ln     string `json:"ln"`
	Email  string `json:"email"`
	Pwd    string `json:"pwd"`
	Jadmin int64  `json:"jadmin"`
}

type Login struct {
	Status   string      `json:"status"`
	Data     UserDetails `json:"data"`
	Login_id int64       `json:"id"`
}

type Login_err struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
