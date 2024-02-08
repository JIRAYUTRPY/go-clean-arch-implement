package services

import (
	"github.com/jirayutrpy/server-go/v2/entities"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/auth"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase interface {
	Register(request entities.User) error
}

type RegisterService struct {
	repo repositorys.RegisterRepository
}

func NewRegisterService(repo repositorys.RegisterRepository) RegisterUseCase{
	return &RegisterService{repo: repo}
}
func (s *RegisterService) Register(request entities.User) error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	request.Password = string(hashedPassword)
	request.Role = "user"
	return s.repo.Save(request)
}

