package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type GetEventRepository interface {
	GetByUserId(userId int) (data []entities.Event, err error)
}