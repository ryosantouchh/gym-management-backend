package models

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
	IsActive  bool   `json:"isActive"`
}
