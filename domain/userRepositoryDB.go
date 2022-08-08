package domain

import (
	"fmt"
	"miniproject_products/errs"
	"miniproject_products/logger"

	"gorm.io/gorm"
)

type UserRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB (client *gorm.DB) UserRepositoryDB{
	return UserRepositoryDB{client}
}

func (s UserRepositoryDB) RegisterUser(users Users)(Users, *errs.AppErr){
	err := s.db.Create(&users).Error
	fmt.Println(err)

	if err != nil {
		logger.Error("Regist User Error")
		return users, errs.NewUnexpectedError("unexpected error")
	}
	return users, nil
}