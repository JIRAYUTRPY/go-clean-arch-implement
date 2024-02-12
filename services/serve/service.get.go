package services

import (
	"errors"

	"github.com/jirayutrpy/server-go/v2/interfaces"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/serve"
)

type GetServeUseCase interface {
	Gets() error
	GetByUserId(id int) (data []interfaces.ServeResponse, err error)
	GetByServeId(id int) (data interfaces.ServeResponse, err error)
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

func (s *GetServeService) GetByServeId(id int) (data interfaces.ServeResponse, err error){
	repsponse, err := s.repo.GetByServeId(id)
	if err != nil {
		return data, err
	}
	data.ID = repsponse.ID
	data.Name = repsponse.Name
	data.Duration = repsponse.Duration
	data.Color = repsponse.Color
	data.UserId = repsponse.UserId
	return data,nil
}
func (s *GetServeService) GetByUserId(id int) (data []interfaces.ServeResponse, err error){
	dataFormat, err := s.repo.GetByUserId(id)
	if err != nil {
		return data, err
	}
	return dataFormat, nil
}
func (s *GetServeService) GetByServeName(name string) error{
	return errors.New("Not yet implement at GetByServeName serve service")
}