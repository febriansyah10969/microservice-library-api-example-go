package dto

type BookHistoriesResponse struct {
	UUID   string `json:"uuid"`
	BookID int    `json:"book_id"`
	Qty    int    `json:"qty"`
	Type   int    `json:"type"`
}
