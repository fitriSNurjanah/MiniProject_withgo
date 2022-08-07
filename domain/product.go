package domain

import "miniproject_products/errs"

type Products struct {
	ID          int `json:"id" db:"id"`
	Merk        string `json:"merk" db:"merk"`
	Description string `json:"description" db:"description"`
}

type ProductRepository interface {
	FindAll() ([]Products, *errs.AppErr)
	// FindByID(int) (Products, *errs.AppErr)

}