package adapters

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/user"
	"gorm.io/gorm"
)

type GormGetByIdRepository struct {
	db *gorm.DB
}

func NewGormGetByIDRepository(db *gorm.DB) repositorys.GetByIdRepository{
	return &GormGetByIdRepository{db: db}
}

func (r *GormGetByIdRepository) GetById(id uint)(data entities.User, err error){
	var dataFormDatabase entities.User
	result := r.db.Where("id = ?",id).First(&dataFormDatabase)
	if result.Error != nil {
		return data,result.Error
	}
	return dataFormDatabase,nil
}