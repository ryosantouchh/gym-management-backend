package entities

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;"`
	Username  string `gorm:"type:varchar(128);unique;not null;"`
	Firstname string `gorm:"type:varchar(128);not null;"`
	Lastname  string `gorm:"type:varchar(128);not null;"`
	Age       int    `gorm:"not null;"`
	Email     string `gorm:"type:varchar(128);unique;not null;"`
	// Booking
	// Credits
	// Orders
	IsActive  bool      `gorm:"not null;default:true;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
