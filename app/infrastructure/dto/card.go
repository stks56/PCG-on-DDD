package dto

type Card struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type Cards []Card
