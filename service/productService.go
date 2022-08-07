package service

import (
	"miniproject_products/domain"
	"miniproject_products/dto"
	"miniproject_products/errs"
)

type ProductService interface {
	GetAllProduct() ([]domain.Products, *errs.AppErr)
	GetProductID(int) (domain.Products, *errs.AppErr)
	CreateProduct(dto.ProductRequest)(domain.Products, *errs.AppErr)
	DeleteProduct(int)(domain.Products, *errs.AppErr)
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

func (s DefaultProductService) GetProductID(id int) (domain.Products, *errs.AppErr) {
	products, err := s.repo.FindByID(id)
	if err != nil {
		return products, err
	}
	return products, nil
}


func (s DefaultProductService)CreateProduct(request dto.ProductRequest)(domain.Products, *errs.AppErr){
	product := domain.Products{}
	product.Merk = request.Merk
	product.Price = request.Price
	product.Description = request.Description

	product, err := s.repo.CreateProduct(product)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (e DefaultProductService) DeleteProduct(id int) (domain.Products, *errs.AppErr) {
	products, err := e.repo.DeleteProduct(id)
	if err != nil {
		return products, err
	}
	return products, nil
}
