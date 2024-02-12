package services

import (
	"github.com/jirayutrpy/server-go/v2/interfaces"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
)

type UpdateServeUseCase interface {
	Update(data interfaces.ServeUpdateRequest) error
}

type UpdateServeService struct {
	repo repositorys.UpdateServeRepository
}

func NewUpdateServeService(repo repositorys.UpdateServeRepository) UpdateServeUseCase{
	return &UpdateServeService{repo: repo}
}

func (s *UpdateServeService) Update(data interfaces.ServeUpdateRequest) error{
	return s.repo.Update(data)
}