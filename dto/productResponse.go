package dto

type ProductResponse struct {
	ID          int    `json:"id" db:"id"`
	Merk        string `json:"merk" db:"merk"`
	Price       int    `json:"price" db:"price"`
	Description string `json:"description" db:"description"`
}
