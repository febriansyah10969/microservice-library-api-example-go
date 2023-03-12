package repository

import (
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) GetTransaction(trans_id string) (model.Transaction, error) {
	transaction := model.Transaction{}

	err := mysqlQB().
		Select("transactions.id", "book_transactions.book_id", "book_transactions.qty").
		From("transactions").
		LeftJoin("book_transactions on book_transactions.trx_id=transactions.id").
		Where(squirrel.Eq{"transactions.uuid": trans_id}).
		Limit(1).
		Scan(&transaction.ID, &transaction.BookTransaction.BookID, &transaction.BookTransaction.Qty)

	if err != nil {
		log.Printf("failed to get data transactions table triggered by void service -> %v", err)
		return model.Transaction{}, errors.New("something wrong happened")
	} else {
		log.Println("success Get User Transaction")
	}

	return transaction, nil
}

func (br *bookRepository) CreateUserTransaction(transaction model.Transaction) (int, error) {

	row, err := mysqlQB().Insert("transactions").Columns("uuid", "code_trx", "user_id", "days", "status", "final_price").
		Values(transaction.UUID, transaction.CodeTrx, transaction.UserID, transaction.Days, transaction.Status, transaction.FinalPrice).Exec()

	if err != nil {
		log.Printf("failed to insert to transactions table triggered by void service -> %v", err)
		return 0, errors.New("something wrong happened")
	} else {
		log.Println("success Create User Transaction")
	}

	insert_id, _ := row.LastInsertId()
	return int(insert_id), nil
}

func (br *bookRepository) CreateBookTransaction(transaction model.BookTransaction) error {
	_, err := mysqlQB().Insert("book_transactions").Columns("trx_id", "uuid", "book_id", "qty", "price").
		Values(transaction.TrxID, transaction.UUID, transaction.BookID, transaction.Qty, transaction.Price).Exec()

	if err != nil {
		log.Printf("failed to insert to book transactions table triggered by void service -> %v", err)
		return errors.New("something wrong happened")
	} else {
		log.Println("success Create Book Transaction")
	}

	return nil
}

func (br *bookRepository) UpdateUserTransaction(trans_id int, transaction model.Transaction) error {
	_, err := mysqlQB().
		Update("transactions").
		Set("days", transaction.Days).
		Set("status", transaction.Status).
		Where(squirrel.Eq{"id": trans_id}).
		Exec()

	if err != nil {
		log.Printf("cannot update Transaction -> Error: %v", err)
		return errors.New("something wrong happened")
	} else {
		log.Printf("Success update Transaction")
	}

	return nil
}

func (br *bookRepository) UpdateBookTransaction(trans_id int, transaction model.BookTransaction) error {
	_, err := mysqlQB().
		Update("book_transactions").
		Set("qty", transaction.Qty).
		Where(squirrel.Eq{"trx_id": trans_id}).
		Exec()

	if err != nil {
		log.Printf("cannot update Transaction Book -> Error: %v", err)
		return errors.New("something wrong happened")
	} else {
		log.Printf("Success update Book Transaction")
	}

	return nil
}
