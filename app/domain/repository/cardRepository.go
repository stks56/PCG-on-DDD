package repository

import "stks56/PCG-onDDD/app/domain/model"

type CardRepository interface {
	GetByID(id int) (*model.Card, error)
}
