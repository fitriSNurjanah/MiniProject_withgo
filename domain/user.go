package domain

import (
	"miniproject_products/errs"

	"github.com/golang-jwt/jwt/v4"
)

type Users struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type AccessTokenClaims struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	jwt.StandardClaims
}


type UserRepository interface {
	RegisterUser(Users) (Users, *errs.AppErr)
}