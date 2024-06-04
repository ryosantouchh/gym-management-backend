package entities

import "time"

type Credit struct {
	ID           uint `gorm:"primaryKey;autoIncrement;"`
	CreditType   uint `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Price        int  `gorm:"not null;"`
	SpecialPrice int  `gorm:"not null;default:0"`

	UserCredit UserCredit `gorm:"foreignKey:CreditID;"`
	CreditUsed CreditUsed `gorm:"foreignKey:CreditID;"`
	OrderItem  OrderItem  `gorm:"foreignKey:CreditID;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
