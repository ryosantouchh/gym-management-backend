package entities

import "time"

type CreditType struct {
	ID            uint   `gorm:"primaryKey;autoIncrement;"`
	Type          string `gorm:"type:varchar(255);not null;unique;"`
	IsMembership  bool   `gorm:"not null;"`
	ClassDuration uint   `gorm:"default:75;"`

	Credit Credit `gorm:"foreignKey:CreditType;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
