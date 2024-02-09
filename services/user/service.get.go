package services

import (
	"errors"

	"github.com/jirayutrpy/server-go/v2/interfaces"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/user"
)

type GetUserUseCase interface {
	Gets()(response []interfaces.GetResponse, err error)
	GetById(id uint) (response interfaces.GetResponse, err error)
	GetByEmail(email string)(response interfaces.GetResponse, err error)
}

type GetUserService struct {
	repo repositorys.GetUserRepository
}

func NewGetUserService(repo repositorys.GetUserRepository) GetUserUseCase{
	return &GetUserService{repo: repo}
}

func (s *GetUserService) Gets()(response []interfaces.GetResponse, err error){
	return response, errors.New("Not yet implement at Gets service")
}

func (s *GetUserService) GetById(id uint) (response interfaces.GetResponse, err error){
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

func (s *GetUserService) GetByEmail(email string)(response interfaces.GetResponse, err error){
	data, err := s.repo.GetByEmail(email)
	if err != nil {
		return response, err
	}
	response.ID = data.ID
	response.Email = email
	response.StoreName = data.StoreName
	response.Address = data.Address
	response.LocationLink = data.LocationLink
	response.Role = data.Role
	return response, err
}

