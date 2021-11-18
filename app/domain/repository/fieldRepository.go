package repository

import "stks56/PCG-on-DDD/app/domain/model"

type FieldRepository interface {
	Get() (*model.Field, error)
}
