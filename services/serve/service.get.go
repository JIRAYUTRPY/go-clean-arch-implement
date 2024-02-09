package services

import (
	"errors"

	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
)

type GetServeUseCase interface {
	Gets() error
	GetByServeId(id uint) error
	GetByUserId(id uint) error
	GetByServeName(name string) error
}
type GetServeService struct {
	repo repositorys.GetServeRepository
}

func NewGetServeService(repo repositorys.GetServeRepository) GetServeUseCase{
	return &GetServeService{repo: repo}
}

func (s *GetServeService) Gets() error{
	return errors.New("Not yet implement at Gets serve service")
}

func (s *GetServeService) GetByServeId(id uint) error{
	return errors.New("Not yet implement at GetByServeId serve service")
}
func (s *GetServeService) GetByUserId(id uint) error{
	return errors.New("Not yet implement at GetByUserId serve service")
}
func (s *GetServeService) GetByServeName(name string) error{
	return errors.New("Not yet implement at GetByServeName serve service")
}