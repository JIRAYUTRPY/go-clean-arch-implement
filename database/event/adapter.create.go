package adapters

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/event"
	"gorm.io/gorm"
)

type GormCreateEventRepository struct {
	db *gorm.DB
}

func NewGormCreateEventRepository(db *gorm.DB) repositorys.CreateEventRepository{
	return &GormCreateEventRepository{db: db}
}

func (r *GormCreateEventRepository) Create(data entities.Event) error{
	return r.db.Create(&data).Error
}