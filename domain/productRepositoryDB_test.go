package domain

import (
	"miniproject_products/dto"
	"miniproject_products/errs"
	"reflect"
	"testing"

	"gorm.io/gorm"
)

func TestNewProductRepositoryDB(t *testing.T) {
	type args struct {
		client *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want ProductRepositoryDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductRepositoryDB(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepositoryDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepositoryDB_FindAll(t *testing.T) {
	type args struct {
		pagination dto.Pagination
	}
	tests := []struct {
		name  string
		s     *ProductRepositoryDB
		args  args
		want  dto.Pagination
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FindAll(tt.args.pagination)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDB.FindAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductRepositoryDB.FindAll() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductRepositoryDB_FindByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		s     ProductRepositoryDB
		args  args
		want  Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.FindByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDB.FindByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductRepositoryDB.FindByID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductRepositoryDB_CreateProduct(t *testing.T) {
	type args struct {
		products Products
	}
	tests := []struct {
		name  string
		s     ProductRepositoryDB
		args  args
		want  Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.CreateProduct(tt.args.products)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDB.CreateProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductRepositoryDB.CreateProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductRepositoryDB_UpdateProduct(t *testing.T) {
	type args struct {
		id       int
		products Products
	}
	tests := []struct {
		name  string
		s     ProductRepositoryDB
		args  args
		want  Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.UpdateProduct(tt.args.id, tt.args.products)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDB.UpdateProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductRepositoryDB.UpdateProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestProductRepositoryDB_DeleteProduct(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name  string
		s     ProductRepositoryDB
		args  args
		want  Products
		want1 *errs.AppErr
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.s.DeleteProduct(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductRepositoryDB.DeleteProduct() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ProductRepositoryDB.DeleteProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
