package domain

import (
	"fmt"
	"miniproject_products/errs"
	"miniproject_products/logger"

	"gorm.io/gorm"
)

type ProductRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB (client *gorm.DB) ProductRepositoryDB{
	return ProductRepositoryDB{client}
}

func (s *ProductRepositoryDB)FindAll()([]Products, *errs.AppErr){
	var products []Products
	err := s.db.Find(&products).Error
	if err != nil {
		logger.Error("error fetch data to customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return products, nil
}

func (s ProductRepositoryDB) FindByID(id int) (Products, *errs.AppErr) {

	var products Products
	err := s.db.First(&products, "id = ?", id)
	if err != nil {
		logger.Error("error fetch data to products table")
		return products, errs.NewUnexpectedError("unexpected error")
	}
	return products, nil
}

func (s ProductRepositoryDB) CreateProduct(products Products)(Products, *errs.AppErr){
	err := s.db.Create(&products).Error
	fmt.Println(err) 

	if err != nil {
		logger.Error("Error created product")
		return products, errs.NewUnexpectedError("unexpected error")
	}
	return products, nil
}

func (s ProductRepositoryDB) DeleteProduct(id int)(Products, *errs.AppErr){
	var products Products
	err := s.db.Delete(&products, "id = ? ", id)
	if err != nil {
		logger.Error("Error Deleting Data")
		return products, errs.NewUnexpectedError("unexpected error")
	}
	return products, nil
}