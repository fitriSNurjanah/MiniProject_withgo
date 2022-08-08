package dto

type UserRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}