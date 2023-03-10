package dto

type TransactionRequest struct {
	BookUUID string `form:"book_uuid" json:"book_uuid" binding:"required"`
	Qty      int    `form:"qty" json:"qty" binding:"required"`
	Days     int    `form:"days" json:"days" binding:"required"`
}
