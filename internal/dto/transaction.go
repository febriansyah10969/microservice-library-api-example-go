package dto

type TransactionRequest struct {
	Qty  int `form:"qty" json:"qty" binding:"required"`
	Days int `form:"days" json:"days" binding:"required"`
}
