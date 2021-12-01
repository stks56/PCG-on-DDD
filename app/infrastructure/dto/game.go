package dto

type Game struct {
	Id     uint `gorm:"primaryKey"`
	Fields []Field
}
