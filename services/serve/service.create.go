package services

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
)

type CreateServeUseCase interface {
	Create(data entities.Serve) error
}

type CreateServeService struct {
	repo repositorys.CreateServeRepository
}

func NewCreateServeService(repo repositorys.CreateServeRepository) CreateServeUseCase{
	return &CreateServeService{repo: repo}
}

func (s *CreateServeService) Create(data entities.Serve) error{
	err := s.repo.Create(data)
	if err != nil {
		return err
	}
	return nil
} 