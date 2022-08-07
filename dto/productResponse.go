package dto

type ProductResponse struct {
	ID          int    `json:"id" db:"id"`
	Merk        string `json:"merk" db:"merk"`
	Description string `json:"description" db:"description"`
}
