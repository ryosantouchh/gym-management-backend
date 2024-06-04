package entities

import "time"

type _Address struct {
	DeliveryFname string
	DeliveryLname string
	PhoneNumber   string
	Street1       string
	Street2       string
	District      string
	SubDistrict   string
	Province      string
	PostalCode    string
}

type UserAddress struct {
	ID     uint `gorm:"primaryKey;autoIncrement;"`
	UserID uint `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	_Address
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
