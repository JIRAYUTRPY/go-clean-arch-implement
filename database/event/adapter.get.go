package adapters

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/event"
	"gorm.io/gorm"
)

type GormGetEventRepository struct {
	db *gorm.DB
}

func NewGormGetEventRepository(db *gorm.DB) repositorys.GetEventRepository{
	return &GormGetEventRepository{db: db}
}

func (r *GormGetEventRepository) GetByUserId(userId int) (data []entities.Event, err error){
	var response []entities.Event
	er := r.db.Find(&response).Where("user_id = ?", userId)
	if er.Error != nil {
		return data, er.Error
	}
	return response, nil
}