package model

// table name for book
const BOOKTRANSACTION string = "book_transactions"

type BookTransaction struct {
	ID     int    `db:"id"`
	TrxID  int    `db:"trx_id"`
	UUID   string `db:"uuid"`
	BookID int    `db:"book_id"`
	Qty    int    `db:"qty"`
	Price  int    `db:"price"`
}

type PartialBookTransaction struct {
	BookID *int `db:"book_id"`
	Qty    *int `db:"qty"`
}
