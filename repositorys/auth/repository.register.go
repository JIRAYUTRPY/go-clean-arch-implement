package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type RegisterRepository interface {
	Save(user entities.User) error
}