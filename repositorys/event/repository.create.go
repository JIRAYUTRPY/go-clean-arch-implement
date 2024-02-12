package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type CreateEventRepository interface {
	Create(data entities.Event) error
}