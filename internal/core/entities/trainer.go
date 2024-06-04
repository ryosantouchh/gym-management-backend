package entities

import "time"

type Trainer struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;"`
	Firstname string    `gorm:"type:varchar(255);not null;"`
	Lastname  string    `gorm:"type:varchar(255);not null;"`
	Age       uint      `gorm:"type:uint;"`
	Classes   []Class   `gorm:"many2many:class_trainers;foreignKey:ID;joinForeignKey:TrainerID;references:ID;joinReferences:ClassID;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}
