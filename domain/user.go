package domain

import "miniproject_products/errs"

type Users struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserRepository interface {
	RegisterUser(Users) (Users, *errs.AppErr)
}