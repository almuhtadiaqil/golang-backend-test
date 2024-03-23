package request

type StoreProductRequest struct {
	ProductName  string  `json:"product_name" binding:"required"`
	ProductCode  string  `json:"product_code" binding:"required"`
	ProductType  string  `json:"product_type" binding:"required"`
	ProductPrice float32 `json:"product_price" binding:"required"`
	ProductQty   int     `json:"product_qty" binding:"required"`
}
