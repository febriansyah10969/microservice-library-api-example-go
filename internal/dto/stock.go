package dto

type StockRequest struct {
	Qty int `form:"qty" json:"qty" binding:"required"`
}
