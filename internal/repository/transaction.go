package repository

import (
	"errors"
	"log"

	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) CreateUserTransaction(transaction model.Transaction) (int, error) {

	row, err := mysqlQB().Insert("transactions").Columns("uuid", "code_trx", "user_id", "days", "status", "final_price").
		Values(transaction.UUID, transaction.CodeTrx, transaction.UserID, transaction.Days, transaction.Status, transaction.FinalPrice).Exec()

	if err != nil {
		log.Printf("failed to insert to transactions table triggered by void service -> %v", err)
		return 0, errors.New("something wrong happened")
	}

	insert_id, _ := row.LastInsertId()
	return int(insert_id), nil
}

func (br *bookRepository) CreateBookTransaction(transaction model.BookTransaction) error {
	_, err := mysqlQB().Insert("transaction_books").Columns("trx_id", "uuid", "book_id", "qty", "price").
		Values(transaction.TrxID, transaction.UUID, transaction.BookID, transaction.Qty, transaction.Price).Exec()

	if err != nil {
		log.Printf("failed to insert to transactions table triggered by void service -> %v", err)
		return errors.New("something wrong happened")
	}

	return nil
}
