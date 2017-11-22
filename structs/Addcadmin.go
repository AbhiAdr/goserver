package structs

type CadminDetail struct {
	Id     int64  `json:"id"`
	Name   string `json:"cname"`
	Url    string `json:"curl"`
	Active int64  `json:"active"`
}

type CadminDetails []CadminDetail

type Addcadmin struct {
	Status string       `json:"status"`
	Data   CadminDetail `json:"data"`
	Id     int64        `json:"id"`
}
