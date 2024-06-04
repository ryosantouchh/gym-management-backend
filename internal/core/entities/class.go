package entities

import "time"

type Class struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;"`
	ClassType string `gorm:"type:varchar(255);unique;not null;"`
	Duration  uint   `gorm:"not null;"`

	Trainers []Trainer `gorm:"many2many:class_trainers;foreignKey:ID;joinForeignKey:ClassID;references:ID;joinReferences:TrainerID;"`
	Booking  Booking   `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Schedule Schedule  `gorm:"foreignKey:ClassID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
