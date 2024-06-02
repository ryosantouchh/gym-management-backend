package entities

type ClassTrainer struct {
	ClassID   uint `gorm:"primaryKey; not null"`
	TrainerID uint `gorm:"primaryKey; not null"`
}
