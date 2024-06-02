package entities

import "time"

type Trainer struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	Firstname string    `gorm:"type:varchar(128);not null;"`
	Lastname  string    `gorm:"type:varchar(128);not null;"`
	Age       int       `gorm:"type:int;"`
	Classes   []Class   `gorm:"many2many:class_trainers;foreignKey:TrainerID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
