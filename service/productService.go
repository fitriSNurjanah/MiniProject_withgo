package service

import (
	"miniproject_products/domain"
	"miniproject_products/errs"
)

type ProductService interface {
	GetAllProduct() ([]domain.Products, *errs.AppErr)
	// GetProductID(int) (domain.Products, *errs.AppErr)
}

type DefaultProductService struct {
	repo domain.ProductRepository
}

func NewProductService(repository domain.ProductRepository) DefaultProductService{
	return DefaultProductService{repository}
}

func (s DefaultProductService) GetAllProduct() ([]domain.Products, *errs.AppErr) {

	products, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

// func (e DefaultProductService) GetProductID(id int) (domain.Products, *errs.AppErr) {
// 	employees, err := e.repo.FindByID(id)
// 	if err != nil {
// 		return employees, err
// 	}
// 	return employees, nil
// }

