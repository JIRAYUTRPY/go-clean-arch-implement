package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type CreateServeRepository interface {
	Create(data entities.Serve) error
}