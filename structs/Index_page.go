package structs

type Data struct {
	PageName string
}

type IndexData struct {
	PageName     string
	LoginId      int64
	Login_status int64
}

type UserDetails struct {
	Id           int64  `json:"id"`
	Fn           string `json:"fn"`
	Ln           string `json:"ln"`
	Email        string `json:"email"`
	Pwd          string `json:"pwd"`
	Login_status int64  `json:"login_status"`
}

type ShowUserDetails []UserDetails
