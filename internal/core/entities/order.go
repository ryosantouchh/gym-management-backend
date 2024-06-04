package entities

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey;autoIncrement;"`
	UserID     uint        `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	TotalPrice uint `gorm:"not null;"`

	DeliveryLname string `gorm:"type:varchar(255);"`
	DeliveryFname string `gorm:"type:varchar(255);"`
	PhoneNumber   string `gorm:"type:varchar(255);"`
	Street1       string `gorm:"type:varchar(255);"`
	Street2       string `gorm:"type:varchar(255);"`
	District      string `gorm:"type:varchar(255);"`
	SubDistrict   string `gorm:"type:varchar(255);"`
	Province      string `gorm:"type:varchar(255);"`
	PostalCode    string `gorm:"type:varchar(255);"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
