package entities

import "time"

type _Address struct {
	DeliveryLname string
	DeliveryFname string
	PhoneNumber   string
	Street1       string
	Street2       string
	District      string
	SubDistrict   string
	Province      string
	PostalCode    string
}

type UserAddress struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;"`
	UserID        uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeliveryLname string    `gorm:"type:varchar(255);"`
	DeliveryFname string    `gorm:"type:varchar(255);"`
	PhoneNumber   string    `gorm:"type:varchar(255);"`
	Street1       string    `gorm:"type:varchar(255);"`
	Street2       string    `gorm:"type:varchar(255);"`
	District      string    `gorm:"type:varchar(255);"`
	SubDistrict   string    `gorm:"type:varchar(255);"`
	Province      string    `gorm:"type:varchar(255);"`
	PostalCode    string    `gorm:"type:varchar(255);"`
	CreatedAt     time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt     time.Time `gorm:"not null;autoUpdateTime;"`
}
