package dto

type ProductRequest struct {
	Merk        string `json:"merk" db:"merk"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}