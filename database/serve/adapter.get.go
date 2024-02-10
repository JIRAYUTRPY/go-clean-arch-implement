package adapters

import (
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
	var dataFormDatabase []entities.Serve
	result := r.db.Limit(10).Find(&dataFormDatabase)
	if result.Error != nil {
		return data, result.Error
	}
	return dataFormDatabase,nil
}

func (r *GormGetServeRepository) GetByServeId(id uint)(data entities.Serve, err error){
	var dataFormDatabase entities.Serve
	result := r.db.Where("id = ?", id).First(&dataFormDatabase)
	if result.Error != nil {
		return data,result.Error
	}
	return data,nil
}
func (r *GormGetServeRepository) GetByServeName(name string)(data []entities.Serve, err error){
	var dataFormDatabase []entities.Serve
	result := r.db.Where("name = ?", name).Limit(10).Find(&dataFormDatabase)
	if result.Error != nil {
		return data, result.Error
	}
	return dataFormDatabase,nil
}

func (r *GormGetServeRepository) GetByUserId(userId uint)(data []entities.Serve, err error){
	var dataFormDatabase []entities.Serve
	result := r.db.Where("user_id = ?", userId).Limit(10).Find(&dataFormDatabase)
	if result.Error != nil {
		return data, result.Error
	}
	return dataFormDatabase,nil
}

