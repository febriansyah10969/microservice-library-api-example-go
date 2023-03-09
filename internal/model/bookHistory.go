package model

// table name for book history
const BOOKHISTORY string = "book_histories"

type BookHistory struct {
	ID     int    `db:"id"`
	UUID   string `db:"uuid"`
	BookID int    `db:"book_id"`
	Qty    int    `db:"qty"`
	Type   int    `db:"type"`
}
