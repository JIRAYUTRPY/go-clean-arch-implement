package adapters

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	"github.com/jirayutrpy/server-go/v2/interfaces"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
	"gorm.io/gorm"
)

type GormUpdateServeRepository struct {
	db *gorm.DB
}

func NewGormUpdateServeRepository(db *gorm.DB) repositorys.UpdateServeRepository{
	return &GormUpdateServeRepository{db: db}
}

func (r *GormUpdateServeRepository) Update(data interfaces.ServeUpdateRequest) error {
	return r.db.Model(entities.Serve{}).Where("ID = ?", data.ID).Updates(entities.Serve{
		Name: data.Name,
		Duration: data.Duration,
		Color: data.Color,
	}).Error
}