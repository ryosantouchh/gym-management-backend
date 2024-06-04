package entities

import "time"

type CreditUsed struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;"`
	UserID       uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserCreditID uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreditID     uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt    time.Time `gorm:"not null;autoUpdateTime;"`
}
