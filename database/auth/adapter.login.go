package adapters

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/auth"
	"gorm.io/gorm"
)

type GormLoginRepository struct {
	db *gorm.DB
}

func NewGormLoginRepository(db *gorm.DB) repositorys.LoginRepository{
	return &GormLoginRepository{db: db}
}

func (r *GormLoginRepository) GetByEmail(value string) (data entities.User, err error){
	var dataFormDatabase entities.User
	result := r.db.Where("email = ?",value).First(&dataFormDatabase)
	if result.Error != nil {
		return dataFormDatabase, result.Error
	}
	return dataFormDatabase, nil
}