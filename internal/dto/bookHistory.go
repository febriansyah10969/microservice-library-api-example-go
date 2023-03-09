package dto

// uuid	varchar(255)
// book_id	bigint unsigned
// qty	int
// type	int

type BookHistoryResponse struct {
	UUID          string          `json:"uuid"`
	Name          string          `json:"name"`
	Stock         int             `json:"stock"`
	BookHistories []BookHistories `json:"book_histories"`
}

type BookHistories struct {
	UUID   string `json:"uuid"`
	BookID int    `json:"book_id"`
	Qty    int    `json:"qty"`
	Type   int    `json:"type"`
}
