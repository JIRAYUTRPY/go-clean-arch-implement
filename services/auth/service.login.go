package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jirayutrpy/server-go/v2/interfaces"
	repositorys "github.com/jirayutrpy/server-go/v2/repositorys/auth"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Login(request interfaces.LoginRequest) (response interfaces.LoginResponse ,userId uint ,err error)
}

type LoginService struct {
	repo repositorys.LoginRepository
}

func NewLoginService(repo repositorys.LoginRepository) LoginUseCase{
	return &LoginService{repo: repo}
}

func (s *LoginService) Login(request interfaces.LoginRequest)(response interfaces.LoginResponse ,userId uint ,err error){
	dataWithPassword, err := s.repo.GetByEmail(request.Email)
	if err != nil {
		return response, userId, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(dataWithPassword.Password),
		[]byte(request.Password),
	)
	if err != nil {
		return response, userId, errors.New("Password Invalid")
	}
	jwtSecretKey := os.Getenv("SECRET_KEY")
	newToken := jwt.New(jwt.SigningMethodHS256)
	claims := newToken.Claims.(jwt.MapClaims)
	claims["id"] = dataWithPassword.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	fmt.Println(time.Now().Add(time.Hour * 72).Unix())
	finishToken, err := newToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return response, userId, err
	}
	response.Token = finishToken
	return response, dataWithPassword.ID, nil
}