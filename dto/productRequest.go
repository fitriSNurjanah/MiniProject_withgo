package dto

type ProductRequest struct {
	Merk        string `json:"merk" db:"merk"`
	Description string `json:"description" db:"description"`
}