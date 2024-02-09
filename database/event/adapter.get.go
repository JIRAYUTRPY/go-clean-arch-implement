package adapters

import (
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/event"
	"gorm.io/gorm"
)

type GormGetEventRepository struct {
	db *gorm.DB
}

func NewGormGetEventRepository(db *gorm.DB) repositorys.GetEventRepository{
	return &GormGetEventRepository{db: db}
}