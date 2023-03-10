package model

// table name for book
const BOOKTRANSACTION string = "transaction_books"

type BookTransaction struct {
	ID     int    `db:"id"`
	TrxID  int    `db:"trx_id"`
	UUID   string `db:"uuid"`
	BookID int    `db:"book_id"`
	Qty    int    `db:"qty"`
	Price  int    `db:"price"`
}
