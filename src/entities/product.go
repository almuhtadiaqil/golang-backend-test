package entities

import (
	"database/sql"
	"time"
)

type Product struct {
	ProductID    uint         `db:"product_id" json:"product_id" gorm:"primaryKey;autoIncrement"`
	ProductName  string       `db:"product_name" json:"product_name" gorm:"not null;"`
	ProductCode  string       `db:"product_code" json:"product_code" gorm:"not null;unique"`
	ProductType  string       `db:"product_type" json:"product_type" gorm:"not null;"`
	ProductPrice float32      `db:"product_price" json:"product_price" gorm:"not null;"`
	ProductQty   int          `db:"product_qty" json:"product_qty" gorm:"not null;"`
	CreatedAt    time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt    sql.NullTime `db:"updated_at" json:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

type ProductResponse struct {
	ProductID    uint      `json:"product_id"`
	ProductName  string    `json:"product_name"`
	ProductCode  string    `json:"product_code"`
	ProductType  string    `json:"product_type"`
	ProductPrice float32   `json:"product_price"`
	ProductQty   int       `json:"product_qty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (p *Product) FormatAsResponse() *ProductResponse {
	return &ProductResponse{
		ProductID:    p.ProductID,
		ProductName:  p.ProductName,
		ProductCode:  p.ProductCode,
		ProductType:  p.ProductType,
		ProductPrice: p.ProductPrice,
		ProductQty:   p.ProductQty,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt.Time,
	}
}
