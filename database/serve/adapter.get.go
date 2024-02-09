package adapters

import (
	"errors"

	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
	"gorm.io/gorm"
)

type GormGetServeRepository struct {
	db *gorm.DB
}

func NewGormGetServeRepository(db *gorm.DB) repositorys.GetServeRepository{
	return &GormGetServeRepository{db: db}
}

///* Have to Implement
func (r *GormGetServeRepository) Gets() (data []entities.Serve, err error){
	return data, errors.New("Not yet implement at Gets Adapter")
}

///* Have to Implement
func (r *GormGetServeRepository) GetByServeId(userId uint)(data entities.Serve, err error){
	return data,errors.New("Not yet implement at GetByServeId Adapter")
}

///* Have to Implement
func (r *GormGetServeRepository) GetByServeName(name string)(data entities.Serve, err error){
	return data,errors.New("Not yet implement at GetByServeName Adapter")
}

///* Have to Implement
func (r *GormGetServeRepository) GetByUserId(id uint)(data entities.Serve, err error){
	return data,errors.New("Not yet implement at GetByUserId Adapter")
}

