package repository

import (
	"errors"
	"log"
	"sort"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
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

func (br *bookRepository) UpdateStock(uuid dto.GetUUID, stock int) error {
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

func (br *bookRepository) GetBookHistory(uuid dto.GetUUID, p *helper.InPage) ([]model.BookHistory, *helper.Pagination, error) {
	bookHistories := []model.BookHistory{}

	sc := mysqlQB().
		Select("COUNT(pr.id)").
		From("book_histories pr").
		LeftJoin("books bh on bh.id = pr.book_id").
		Where(squirrel.Eq{"pr.uuid": uuid.UUID})

	qb := mysqlQB().
		Select("pr.id", "pr.uuid", "pr.qty", "pr.type").
		From("book_histories pr").
		LeftJoin("books b on b.id = pr.book_id").
		Where(squirrel.Eq{"b.uuid": uuid.UUID})

	qb, pag := Paginate(sc, qb, *p)

	rows, err := qb.Query()
	if err != nil {
		log.Printf("Query rows failed: %v", err)
		return bookHistories, pag, errors.New("something wrong happened")
	}

	for rows.Next() {
		bookHistory := model.BookHistory{}
		if err := rows.Scan(&bookHistory.BookID, &bookHistory.UUID, &bookHistory.Qty, &bookHistory.Type); err != nil {
			log.Printf("scan rows failed: %v", err)
			return bookHistories, pag, errors.New("something wrong happened")
		}

		bookHistories = append(bookHistories, bookHistory)
	}

	defer rows.Close()

	if len(p.Perpage) == 0 {
		defaultPerPage := "10"
		pag.Perpage = &defaultPerPage
	} else {
		pag.Perpage = &p.Perpage
	}

	sort.SliceStable(bookHistories, func(i, j int) bool {
		return bookHistories[i].ID > bookHistories[j].ID
	})

	var hmp bool
	if len(bookHistories) > 0 {
		more := sc.Where(squirrel.Lt{"id": bookHistories[len(bookHistories)-1].ID})
		hmp = hasMorePages(more)
	} else {
		hmp = false
	}

	pag.HasMorePages = &hmp
	if hmp {
		first := curs{bookHistories[len(bookHistories)-1].ID, *pag.HasMorePages}
		enc := encodeCursor(first)
		pag.NextCursor = &enc
	} else {
		pag.NextCursor = nil
	}

	if len(bookHistories) > 0 {
		sc = sc.Where(squirrel.Gt{"id": bookHistories[0].ID})
		if isFirstPage(sc) {
			pag.PrevCursor = nil
		} else {
			c := curs{bookHistories[0].ID, false}
			enc := encodeCursor(c)
			pag.PrevCursor = &enc
		}
	} else {
		pag.PrevCursor = nil
	}

	return bookHistories, pag, nil
}
