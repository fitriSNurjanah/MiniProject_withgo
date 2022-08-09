package domain

import (
	"fmt"
	"miniproject_products/dto"
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

func (s *ProductRepositoryDB)FindAll(pagination dto.Pagination)(dto.Pagination, *errs.AppErr){
	var p dto.Pagination
	tr := 0
	offset := pagination.Page  * pagination.Limit


	var products []Products
	err :=s.db.Limit(pagination.Limit).Offset(offset).Find(&products).Error
	if err != nil {
		return p, nil
	}
	pagination.Rows = products
	totalRows := int64(tr)

	errCount := s.db.Model(products).Count(&totalRows).Error

	if errCount!= nil{
		return p, errs.NewUnexpectedError("unexpected error")
	}
	return pagination, nil
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


func (s ProductRepositoryDB)UpdateProduct(id int, products Products)(Products, *errs.AppErr){
	err := s.db.Model(&products).Where("id = ? ", id).Updates(products)

	if err != nil{
		logger.Info("updating data errors")
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