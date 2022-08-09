package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	GenerateToken(int) (string, error)
	ValidateToken(string) (string, bool, error)
}


type JwtService struct {
	UserID string `json:"id"`
	jwt.RegisteredClaims
}

func NewAuthService() *JwtService{
	return &JwtService{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 *time.Hour)),
		},
	}
}

func (s *JwtService) GenerateToken(userID int)(string, error){
	claim := NewAuthService()
	claim.UserID = strconv.Itoa(userID)
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte("RAHASIA"))
		if err != nil {
			fmt.Println("error : ", err)
			return tokenString, err
		}
		return tokenString, nil
}

func (s *JwtService) ValidateToken(encodedToken string)(string, bool, error){
	
	tokenString, err := jwt.ParseWithClaims(encodedToken, &JwtService{}, func(token *jwt.Token)(interface{}, error){
		return []byte("RAHASIA"), nil
	})

	if err != nil {
		fmt.Println("errro : ", err.Error())
		return "0", false, err
	}

	if claim, ok := tokenString.Claims.(*JwtService); ok &&tokenString.Valid{
		fmt.Println("Test : ", claim.UserID)
		return claim.UserID, true, nil
	}else {
		return claim.UserID, false, nil
	}
}

