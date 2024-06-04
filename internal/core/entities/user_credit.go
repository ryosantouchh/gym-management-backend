package entities

import "time"

type UserCredit struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	UserID    uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreditID  uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Amount    uint      `gorm:"not null;default:0;"`
	ExpiredAt time.Time `gorm:"not null;"`

	CreditUsed CreditUsed `gorm:"foreignKey:UserCreditID;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
