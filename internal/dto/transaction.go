package dto

type TransactionRequest struct {
	BookUUID string `form:"book_uuid" json:"book_uuid" binding:"required"`
	TransID  string `form:"transaction_id" json:"transaction_id"`
	Qty      int    `form:"qty" json:"qty" binding:"required"`
	Days     int    `form:"days" json:"days" binding:"required"`
}

type TransactionUUIDRequest struct {
	TransUUID string `form:"transaction_id" json:"transaction_id" binding:"required"`
}

type TransactionResponse struct {
	UUID       string `json:"uuid"`
	CodeTrx    string `json:"code_trx"`
	Days       int    `json:"days"`
	Status     int    `json:"status"`
	FinalPrice int    `json:"final_price"`
	BookID     *int   `json:"book_uuid"`
	Qty        *int   `json:"qty"`
}
