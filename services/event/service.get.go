package services

import (
	"errors"

	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/event"
)

type GetEventUsecase interface{
	Gets()(data entities.Event, err error)
	GetById(id uint)(data entities.Event, err error)
	GetByUserId(userId int)(data []entities.Event, err error)
}

type GetEventService struct {
	repo repositorys.GetEventRepository
}

func NewGetEventService(repo repositorys.GetEventRepository) GetEventUsecase{
	return &GetEventService{repo: repo}
}

func (s *GetEventService) Gets()(data entities.Event, err error){
	return data,errors.New("Not yet implement")
}

func (s *GetEventService) GetById(id uint)(data entities.Event, err error){
	return data,errors.New("Not yet implement")
}
func (s *GetEventService) GetByUserId(userId int)(data []entities.Event, err error){
	reponse, err := s.repo.GetByUserId(userId)
	if err != nil {
		return data, err
	}
	return reponse,nil
}