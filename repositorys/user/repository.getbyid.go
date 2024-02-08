package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type GetByIdRepository interface {
	GetById(id uint) (data entities.User, err error)
}