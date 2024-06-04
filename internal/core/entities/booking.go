package entities

import "time"

type Booking struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;"`
	UserID     uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ScheduleID uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ClassID    uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	BookedAt   time.Time `gorm:"not null;"`
	CreatedAt  time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt  time.Time `gorm:"not null;autoUpdateTime;"`
}
