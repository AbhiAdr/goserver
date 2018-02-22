package structs

type UserDetails struct {
	Id           int64  `json:"id"`
	Fn           string `json:"fn"`
	Ln           string `json:"ln"`
	Email        string `json:"email"`
	Pwd          string `json:"pwd"`
	Login_status int64  `json:"login_status"`
}

type ShowUserDetails []UserDetails
