package dto

import "stks56/PCG-on-DDD/app/domain/model"

type Game struct {
	Id     uint `gorm:"primaryKey"`
	Fields []Field
	Status uint8 `gorm:"not null"`
}

func (game *Game) ConvertToModel() *model.Game {
	return &model.Game{
		YourField:     game.Fields[0].ConvertToModel(),
		OpponentField: game.Fields[1].ConvertToModel(),
	}
}
