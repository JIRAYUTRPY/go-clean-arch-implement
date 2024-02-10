package repositorys

import "github.com/jirayutrpy/server-go/v2/entities"

type GetServeRepository interface {
	Gets() (data []entities.Serve, err error)
	GetByServeId(id uint)(data entities.Serve, err error)
	GetByServeName(name string)(data []entities.Serve, err error)
	GetByUserId(userId uint)(data []entities.Serve, err error)
}