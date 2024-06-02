package entities

import "time"

type ClassType string

const (
	Boxing   ClassType = "boxing"
	Muaythai ClassType = "muaythai"
	BJJ      ClassType = "bjj"
)

type Class struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	Type      ClassType `gorm:"type:varchar(128);unique;not null;"`
	Trainers  []Trainer `gorm:"many2many:class_trainers;foreignKey:ClassID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
