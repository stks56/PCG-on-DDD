package repository

import "stks56/PCG-on-DDD/app/domain/model"

type GameRepository interface {
	Insert(game *model.Game) error
}
