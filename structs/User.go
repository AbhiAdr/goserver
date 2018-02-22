package structs

type User struct {
	Id           string `json:"id"`
	Firstname    string `json:"fn"`
	Lastname     string `json:"ln"`
	Email        string `json:"email"`
	Pwd          string `json:"pwd"`
	Login_status string `json:"login_status"`
}

type Users []User
