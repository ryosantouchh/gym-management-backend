package entities

import "time"

type Booking struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;"`
	UserID     uint      `gorm:"foreignKey:UserID;references:ID"`
	ScheduleID uint      `gorm:"foreignKey:ScheduleID;references:ID"`
	BookedAt   time.Time `gorm:"not null;"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;"`
}
