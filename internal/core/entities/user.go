package entities

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;"`
	Username  string `gorm:"type:varchar(255);unique;not null;"`
	Password  string `gorm:"type:varchar(255);not null;"`
	Firstname string `gorm:"type:varchar(255);not null;"`
	Lastname  string `gorm:"type:varchar(255);not null;"`
	Age       uint   `gorm:"not null;"`
	Email     string `gorm:"type:varchar(255);unique;not null;"`
	IsActive  bool   `gorm:"not null;default:true;"`

	OauthProvider string `gorm:"type:varchar(255);"`
	OauthUID      string `gorm:"type:varchar(255);"`
	RefreshToken  string `gorm:"type:varchar(255);"` // optional consider it further + use both oauth and credential login

	Bookings      []Booking     `gorm:"foreignKey:UserID;"`
	UserCredits   []UserCredit  `gorm:"foreignKey:UserID;"`
	CreditUsed    []CreditUsed  `gorm:"foreignKey:UserID;"`
	Orders        []Order       `gorm:"foreignKey:UserID;"`
	UserAddresses []UserAddress `gorm:"foreignKey:UserID;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
