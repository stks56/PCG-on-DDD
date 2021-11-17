package repository

import "stks56/PCG-on-DDD/app/domain/model"

type CardRepository interface {
	GetByID(id int) (*model.Card, error)
}
