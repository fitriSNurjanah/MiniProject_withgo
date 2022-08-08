package app

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestProductHandler_getAllProduct(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *ProductHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.getAllProduct(tt.args.c)
		})
	}
}

func TestProductHandler_getProductID(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *ProductHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.getProductID(tt.args.c)
		})
	}
}

func TestProductHandler_createProduct(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *ProductHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.createProduct(tt.args.c)
		})
	}
}

func TestProductHandler_UpdateProduct(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *ProductHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.UpdateProduct(tt.args.c)
		})
	}
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		ch   *ProductHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ch.DeleteProduct(tt.args.c)
		})
	}
}
