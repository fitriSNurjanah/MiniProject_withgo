package service

import (
	"fmt"
	"miniproject_products/domain"
	"miniproject_products/dto"
	"miniproject_products/errs"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(dto.UserRequest)(domain.Users, *errs.AppErr)
	LoginUser(dto.Login) (domain.Users, *errs.AppErr)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repository domain.UserRepository)DefaultUserService{
	return DefaultUserService{repository}
}

func (s DefaultUserService)RegisterUser(request dto.UserRequest)(domain.Users, *errs.AppErr){
	user := domain.Users{}
	user.Username = request.Username
	
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil{
		return user, errs.NewUnexpectedError("unexpected error")
	}
	fmt.Println("hash password")

	user.Password = string(hashPassword)

	user, errs := s.repo.RegisterUser(user)
	if errs != nil {
		return user, errs
	}
	return user, nil

}


func (s DefaultUserService)LoginUser(request dto.Login)(domain.Users, *errs.AppErr){
	Userneme := request.Username
	Password := request.Password

	user, err := s.repo.LoginUserInput(Userneme)
	if err != nil {
		return user, err
	}

	if Userneme == ""{
		return user, err
	}

	errByc := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Password))
	if errByc != nil{
		return user, err
	}
	return user, nil
}