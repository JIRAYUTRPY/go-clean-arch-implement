package repositorys

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	"github.com/jirayutrpy/server-go/v2/interfaces"
)

type GetServeRepository interface {
	Gets() (data []entities.Serve, err error)
	GetByServeId(id int)(data entities.Serve, err error)
	GetByServeName(name string)(data []entities.Serve, err error)
	GetByUserId(userId int)(data []interfaces.ServeResponse, err error)
}