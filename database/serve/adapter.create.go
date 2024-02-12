package adapters

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
	"gorm.io/gorm"
)

type GormCreateServeRepository struct {
	db *gorm.DB
}

func NewGormCreateServeRepository(db *gorm.DB) repositorys.CreateServeRepository {
	return &GormCreateServeRepository{db: db}
}

func (r *GormCreateServeRepository) Create(data entities.Serve) error {
	return r.db.Create(&data).Error
}