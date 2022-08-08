package service

import (
	"miniproject_products/domain"
	"miniproject_products/dto"
	"miniproject_products/errs"
	"reflect"
	"testing"
)

func TestNewProductService(t *testing.T) {
	type args struct {
		repository domain.ProductRepository
	}
	tests := []struct {
		name string
		args args
		want DefaultProductService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultProductService_GetAllProduct(t *testing.T) {
	type args struct {
		p dto.Pagination
	}
	tests := []struct {
		name  string
		s     DefaultProductService
		args  args
		want  dto.Pagination
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetAllProduct(tt.args.p)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProductService.GetAllProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultProductService.GetAllProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultProductService_GetProductID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		s     DefaultProductService
		args  args
		want  domain.Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.GetProductID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProductService.GetProductID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultProductService.GetProductID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultProductService_CreateProduct(t *testing.T) {
	type args struct {
		request dto.ProductRequest
	}
	tests := []struct {
		name  string
		s     DefaultProductService
		args  args
		want  domain.Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateProduct(tt.args.request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProductService.CreateProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultProductService.CreateProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultProductService_UpdateProduct(t *testing.T) {
	type args struct {
		id      int
		request dto.ProductRequest
	}
	tests := []struct {
		name  string
		s     DefaultProductService
		args  args
		want  domain.Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateProduct(tt.args.id, tt.args.request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProductService.UpdateProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultProductService.UpdateProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDefaultProductService_DeleteProduct(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		s     DefaultProductService
		args  args
		want  domain.Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DeleteProduct(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultProductService.DeleteProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("DefaultProductService.DeleteProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
