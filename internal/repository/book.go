package repository

import (
	"errors"
	"log"

	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) CreateBook(bm model.Book) error {
	var err error

	_, err = mysqlQB().Insert("books").Columns("uuid", "author_id", "name", "price", "stock").
		Values(bm.UUID, bm.AuthorID, bm.Name, bm.Price, bm.Stock).Exec()

	if err != nil {
		log.Printf("failed to insert to recap_cash_voids table triggered by void service -> %v", err)
		return errors.New("something wrong happened")
	}

	return nil
}
