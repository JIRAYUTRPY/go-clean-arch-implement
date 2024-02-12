package adapters

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	"github.com/jirayutrpy/server-go/v2/interfaces"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
	"gorm.io/gorm"
)

type GormGetServeRepository struct {
	db *gorm.DB
}

func NewGormGetServeRepository(db *gorm.DB) repositorys.GetServeRepository{
	return &GormGetServeRepository{db: db}
}

func (r *GormGetServeRepository) Gets() (data []entities.Serve, err error){
	var dataFormDatabase []entities.Serve
	result := r.db.Limit(10).Find(&dataFormDatabase)
	if result.Error != nil {
		return data, result.Error
	}
	return dataFormDatabase,nil
}

func (r *GormGetServeRepository) GetByServeId(id int)(data entities.Serve, err error){
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

func (r *GormGetServeRepository) GetByUserId(userId int)(data []interfaces.ServeResponse, err error){
	var dataFormDatabase []entities.Serve
	var reponse []interfaces.ServeResponse
	result := r.db.Where("user_id = ?", userId).Find(&dataFormDatabase)
	if result.Error != nil {
		return data, result.Error
	}
	for _, data := range dataFormDatabase {
		newData := interfaces.ServeResponse{
			ID: data.ID,
			Name: data.Name,
			Duration: data.Duration,
			Color: data.Color,
		}
		reponse = append(reponse, newData)
	}
	return reponse,nil
}

