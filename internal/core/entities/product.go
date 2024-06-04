package entities

import "time"

type Product struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;"`
	Name         string    `gorm:"not null;"`
	Price        int       `gorm:"not null;"`
	SpecialPrice int       `gorm:"not null;default:0"`
	ExpiredAt    time.Time `gorm:"not null;"`

	OrderItem OrderItem `gorm:"foreignKey:ProductID;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
