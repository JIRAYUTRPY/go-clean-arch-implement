package repositorys

import "github.com/jirayutrpy/server-go/v2/interfaces"

type UpdateServeRepository interface {
	Update(data interfaces.ServeUpdateRequest) error
}