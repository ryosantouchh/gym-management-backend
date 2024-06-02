package domain

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username" gorm:"unique"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Email     string `json:"email" gorm:"unique"`
	IsActive  bool   `json:"isActive"`
}
