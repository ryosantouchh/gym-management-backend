package entities

import "time"

type Credit struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;"`
	CreditTypeID uint      `gorm:"foreignKey:CreditTypeID,references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount       int       `gorm:"not null;default:0;"`
	ExpiredAt    time.Time `gorm:"not null;"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;"`
}
