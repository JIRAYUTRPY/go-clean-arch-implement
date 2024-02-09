package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type GetUserRepository interface {
	Gets() (data []entities.User, err error)
	GetById(id uint) (data entities.User, err error)
	GetByEmail(email string)(data entities.User, err error)
}