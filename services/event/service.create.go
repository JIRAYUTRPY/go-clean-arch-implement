package services

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/event"
)

type CreateEventUseCase interface {
	Create(data entities.Event) error
}

type CreateEventService struct{
	repo repositorys.CreateEventRepository
}

func NewCreateEventService(repo repositorys.CreateEventRepository) CreateEventUseCase{
	return &CreateEventService{repo: repo}
}

func (s *CreateEventService) Create(data entities.Event) error{
	return s.repo.Create(data)
}