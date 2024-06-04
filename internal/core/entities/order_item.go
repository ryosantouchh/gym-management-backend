package entities

import "time"

type OrderItem struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;"`
	OrderID    uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID  uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreditID   uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount     int       `gorm:"not null;"`
	PriceEach  int       `gorm:"not null;"`
	TotalPrice int       `gorm:"not null;"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;"`
}
