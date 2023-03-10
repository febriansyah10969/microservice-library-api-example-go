package model

// table name for book
const TRANSACTION string = "transactions"

type Transaction struct {
	ID              int             `db:"id"`
	UUID            string          `db:"uuid"`
	CodeTrx         string          `db:"code_trx"`
	UserID          int             `db:"user_id"`
	Days            int             `db:"days"`
	Status          int             `db:"status"`
	FinalPrice      int             `db:"final_price"`
	BookTransaction BookTransaction `db:"book_transaction"`
}
