package entities

import "time"

type CreditTypeEnum string

const (
	SingleClass   CreditTypeEnum = "single"
	PrivateClass  CreditTypeEnum = "private"
	MonthlyMember CreditTypeEnum = "monthly"
)

type CreditType struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;"`
	Type      CreditTypeEnum `gorm:"type:varchar(128);unique;not null;"`
	CreatedAt time.Time      `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time      `gorm:"not null;autoUpdateTime;"`
}
