package entities

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey;autoIncrement;"`
	UserID     uint        `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	_Address
	TotalPrice uint      `gorm:"not null;"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;"`
}
