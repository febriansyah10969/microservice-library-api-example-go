package repository

import (
	"errors"
	"log"
	"sort"
	"strconv"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) GetTransactions(f *helper.Filter, p *helper.InPage) ([]model.Transaction, *helper.Pagination, error) {
	transactions := []model.Transaction{}

	sc := mysqlQB().Select("COUNT(id)").From("transactions")

	qb := mysqlQB().
		Select("pr.uuid", "pr.code_trx", "pr.days", "pr.status", "pr.final_price", "book_transactions.book_id", "book_transactions.qty").
		From("transactions pr").
		LeftJoin("book_transactions on book_transactions.trx_id=pr.id")

	qb, pag := Paginate(sc, qb, *p)

	rows, err := qb.Query()

	if err != nil {
		log.Printf("Query Transactions rows failed: %v", err)
		return transactions, pag, errors.New("something wrong happened")
	} else {
		log.Println("success Get Transactions")
	}

	for rows.Next() {
		transaction := model.Transaction{}
		if err := rows.Scan(&transaction.UUID, &transaction.CodeTrx, &transaction.Days, &transaction.Status, &transaction.FinalPrice, &transaction.PartialBookTransaction.BookID, &transaction.PartialBookTransaction.Qty); err != nil {
			log.Printf("scan rows failed: %v", err)
			return transactions, nil, errors.New("something wrong happened")
		} else {
			log.Println("Berhasil mendapatkan data transaksi : " + strconv.Itoa(len(transactions)))
		}

		transactions = append(transactions, transaction)
	}

	if len(p.Perpage) == 0 {
		defaultPerPage := "10"
		pag.Perpage = &defaultPerPage
	} else {
		pag.Perpage = &p.Perpage
	}

	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].ID > transactions[j].ID
	})

	var hmp bool
	if len(transactions) > 0 {
		more := sc.Where(squirrel.Lt{"id": transactions[len(transactions)-1].ID})
		hmp = hasMorePages(more)
	} else {
		hmp = false
	}

	pag.HasMorePages = &hmp
	if hmp {
		first := curs{transactions[len(transactions)-1].ID, *pag.HasMorePages}
		enc := encodeCursor(first)
		pag.NextCursor = &enc
	} else {
		pag.NextCursor = nil
	}

	if len(transactions) > 0 {
		sc = sc.Where(squirrel.Gt{"id": transactions[0].ID})
		if isFirstPage(sc) {
			pag.PrevCursor = nil
		} else {
			c := curs{transactions[0].ID, false}
			enc := encodeCursor(c)
			pag.PrevCursor = &enc
		}
	} else {
		pag.PrevCursor = nil
	}

	return transactions, pag, nil
}

func (br *bookRepository) GetTransaction(trans_id string) (model.Transaction, error) {
	transaction := model.Transaction{}

	err := mysqlQB().
		Select("transactions.id", "transaction.status", "book_transactions.book_id", "book_transactions.qty").
		From("transactions").
		LeftJoin("book_transactions on book_transactions.trx_id=transactions.id").
		Where(squirrel.Eq{"transactions.uuid": trans_id}).
		Limit(1).
		Scan(&transaction.ID, &transaction.Status, &transaction.BookTransaction.BookID, &transaction.BookTransaction.Qty)

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
