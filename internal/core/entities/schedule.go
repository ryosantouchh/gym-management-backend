package entities

import "time"

type Schedule struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	ClassID   uint      `gorm:"foreignKey:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Date      time.Time `gorm:"not null;"`
	StartTime time.Time `gorm:"not null;"`
	EndTime   time.Time `gorm:"not null;"`

	Bookings []Booking `gorm:"foreignKey:ScheduleID;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
