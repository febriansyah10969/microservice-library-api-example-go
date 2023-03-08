package repository

import (
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
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

func (br *bookRepository) UpdateBook(uuid dto.GetUUID, bm model.Book) error {
	_, err := mysqlQB().Update("books").Set("author_id", bm.AuthorID).Set("name", bm.Name).Set("price", bm.Price).Where(squirrel.Eq{"uuid": uuid}).Exec()

	if err != nil {
		log.Printf("cannot update stock -> Error: %v", err)
		return errors.New("something wrong happened")
	} else {
		log.Printf("Success update Books")
	}

	return nil
}
