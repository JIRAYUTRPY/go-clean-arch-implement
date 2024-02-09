package adapters

import (
	"errors"

	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/user"
	"gorm.io/gorm"
)

type GormGetUserRepository struct {
	db *gorm.DB
}

func NewGormGetUserRepository(db *gorm.DB) repositorys.GetUserRepository{
	return &GormGetUserRepository{db: db}
}

///* Have to Implement
func (r *GormGetUserRepository) Gets() (data []entities.User, err error){
	return data, errors.New("Not yet implement at Gets Adapter")
}

func (r *GormGetUserRepository) GetById(id uint)(data entities.User, err error){
	var dataFormDatabase entities.User
	result := r.db.Where("id = ?",id).First(&dataFormDatabase)
	if result.Error != nil {
		return data,result.Error
	}
	return dataFormDatabase,nil
}

///* Have to Implement
func (r *GormGetUserRepository) GetByEmail(email string)(data entities.User, err error){
	return data, errors.New("Not yet implement at GetByEmail Adapter")
}