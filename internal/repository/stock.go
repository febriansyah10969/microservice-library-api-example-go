package repository

import (
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) GetCurrentStock(uuid dto.GetUUID) (model.Book, error) {
	book := model.Book{}

	err := mysqlQB().
		Select("pr.id", "pr.uuid", "pr.author_id", "pr.stock").
		From("books pr").
		Where(squirrel.Eq{"uuid": uuid.UUID}).
		Limit(1).
		Scan(&book.ID, &book.UUID, &book.AuthorID, &book.Stock)

	if err != nil {
		log.Printf("cannot Get stock -> Error: %v", err)
		return model.Book{}, errors.New("something wrong happened")
	} else {
		log.Printf("Success Get Stock")
	}

	return book, nil
}

func (br *bookRepository) IncreaseStock(uuid dto.GetUUID, stock int) error {
	_, err := mysqlQB().Update("books").Set("stock", stock).Where(squirrel.Eq{"uuid": uuid.UUID}).Exec()

	if err != nil {
		log.Printf("cannot update stock -> Error: %v", err)
		return errors.New("something wrong happened")
	} else {
		log.Printf("Success update Stock")
	}

	return nil
}

func (br *bookRepository) CreateBookHistory(mbh model.BookHistory) error {
	var err error

	_, err = mysqlQB().Insert("book_histories").Columns("uuid", "book_id", "qty", "type").
		Values(mbh.UUID, mbh.BookID, mbh.Qty, mbh.Type).Exec()

	if err != nil {
		log.Printf("failed to insert to book_histories table triggered by void service -> %v", err)
		return errors.New("something wrong happened")
	}

	return nil
}
