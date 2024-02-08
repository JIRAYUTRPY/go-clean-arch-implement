package adapters

import (
	"errors"

	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/auth"
	"gorm.io/gorm"
)

type GormRegisterRepository struct {
	db *gorm.DB
}

func NewGormRegisterRepository(db *gorm.DB) repositorys.RegisterRepository{
	return &GormRegisterRepository{db: db}
}

func (r *GormRegisterRepository) Save(user entities.User) error {
	var data entities.User
	if err := r.db.Where("email = ?", user.Email).First(&data).Error; err != nil{
		return r.db.Create(&user).Error
	}
	return errors.New("Email already exist")
}