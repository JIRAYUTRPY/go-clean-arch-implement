package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type LoginRepository interface {
	GetByEmail(value string) (data entities.User, err error)
}