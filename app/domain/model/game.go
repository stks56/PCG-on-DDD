package model

import "stks56/PCG-on-DDD/app/core/enum"

type Game struct {
	YourField     *Field
	OpponentField *Field
	Status        enum.GameStatusEnum
}
