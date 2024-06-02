package entities

import "time"

type Schedule struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	ClassID   uint      `gorm:"foreignKey:ClassID;references:ID"`
	Date      time.Time `gorm:"not null;"`
	StartTime time.Time `gorm:"not null;"`
	EndTime   time.Time `gorm:"not null;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
