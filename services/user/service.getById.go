package services

import (
	"github.com/jirayutrpy/server-go/v2/interfaces"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/user"
)

type GetByIdUseCase interface {
	GetById(id uint) (response interfaces.GetByIDResponse, err error)
}

type GetByIdService struct {
	repo repositorys.GetByIdRepository
}

func NewGetByIdService(repo repositorys.GetByIdRepository) GetByIdUseCase{
	return &GetByIdService{repo: repo}
}

func (s *GetByIdService) GetById(id uint) (response interfaces.GetByIDResponse, err error){
	data, err := s.repo.GetById(id)
	if err != nil {
		return response, err
	}
	response.ID = id
	response.Email = data.Email
	response.StoreName = data.StoreName
	response.Address = data.Address
	response.LocationLink = data.LocationLink
	response.Role = data.Role
	return response, err
}