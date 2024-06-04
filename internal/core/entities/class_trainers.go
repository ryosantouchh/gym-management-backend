package entities

type ClassTrainer struct {
	ClassID   uint `gorm:"primaryKey;"`
	TrainerID uint `gorm:"primaryKey;"`
}
