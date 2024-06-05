package entities

import "time"

type Credit struct {
	ID             uint   `gorm:"primaryKey;autoIncrement;"`
	Type           string `gorm:"type:varchar(255);not null;unique;"`
	IsMembership   bool   `gorm:"not null;"`
	ExpireDuration uint   `gorm:"not null;"`
	Price          int    `gorm:"not null;"`
	SpecialPrice   int    `gorm:"not null;default:0"`

	UserCredit UserCredit `gorm:"foreignKey:CreditID;"`
	CreditUsed CreditUsed `gorm:"foreignKey:CreditID;"`
	OrderItem  OrderItem  `gorm:"foreignKey:CreditID;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
