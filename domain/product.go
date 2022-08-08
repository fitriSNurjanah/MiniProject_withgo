package domain

import "miniproject_products/errs"

type Products struct {
	ID          int `json:"id" db:"id"`
	Merk        string `json:"merk" db:"merk"`
	Price        int `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}

type ProductRepository interface {
	FindAll() ([]Products, *errs.AppErr)
	FindByID(int) (Products, *errs.AppErr)
	CreateProduct(Products) (Products, *errs.AppErr)
	UpdateProduct(int, Products) (Products, *errs.AppErr)
	DeleteProduct(int) (Products, *errs.AppErr)
}