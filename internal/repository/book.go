package repository

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) GetBook(book_uuid dto.GetUUID) (model.Book, error) {
	book := model.Book{}

	err := mysqlQB().
		Select("pr.id", "pr.uuid", "pr.author_id", "pr.name", "pr.price", "pr.stock").
		From("books pr").
		Where(squirrel.Eq{"uuid": book_uuid.UUID}).
		Limit(1).
		Scan(&book.ID, &book.UUID, &book.AuthorID, &book.Name, &book.Price, &book.Stock)

	if err != nil {
		log.Printf("cannot Get book -> Error: %v", err)
		return model.Book{}, errors.New("something wrong happened")
	} else {
		log.Printf("Success Get book")
	}

	return book, nil
}

func (br *bookRepository) GetBookByID(id int) (model.Book, error) {
	book := model.Book{}

	err := mysqlQB().
		Select("pr.id", "pr.uuid", "pr.author_id", "pr.name", "pr.price", "pr.stock").
		From("books pr").
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		Scan(&book.ID, &book.UUID, &book.AuthorID, &book.Name, &book.Price, &book.Stock)

	if err != nil {
		log.Printf("cannot Get book -> Error: %v", err)
		return model.Book{}, errors.New("something wrong happened")
	} else {
		log.Printf("Success Get book")
	}

	return book, nil
}

func (br *bookRepository) GetBooks(f *helper.Filter, p *helper.InPage) ([]model.Book, *helper.Pagination, error) {
	var result []model.Book

	sc := mysqlQB().Select("COUNT(id)").From("books pr")

	qb := mysqlQB().
		Select("pr.id", "pr.uuid", "pr.author_id", "pr.name", "pr.price", "pr.stock").
		From("books pr")

	if len(f.BookUUID) != 0 {
		qb = qb.Where("pr.uuid LIKE ?", "%"+f.BookUUID+"%")
		sc = sc.Where("pr.uuid LIKE ?", "%"+f.BookUUID+"%")
	}

	if f.BookID != 0 {
		qb = qb.Where("pr.id LIKE ?", "%"+strconv.Itoa(f.BookID)+"%")
		sc = sc.Where("pr.id LIKE ?", "%"+strconv.Itoa(f.BookID)+"%")
	}

	if f.AuthorID != 0 {
		qb = qb.Where("pr.author_id LIKE ?", "%"+strconv.Itoa(f.AuthorID)+"%")
		sc = sc.Where("pr.author_id LIKE ?", "%"+strconv.Itoa(f.AuthorID)+"%")
	}

	if len(f.Name) != 0 {
		qb = qb.Where("pr.name LIKE ?", "%"+f.Name+"%")
		sc = sc.Where("pr.name LIKE ?", "%"+f.Name+"%")
	}

	if len(strconv.Itoa(f.MinPrice)) != 1 && len(strconv.Itoa(f.MaxPrice)) == 1 {
		fmt.Println("test1")
		qb = qb.Where("pr.price >= ? ", strconv.Itoa(f.MinPrice))
		sc = sc.Where("pr.price >= ? ", strconv.Itoa(f.MinPrice))
	} else if len(strconv.Itoa(f.MinPrice)) == 1 && len(strconv.Itoa(f.MaxPrice)) != 1 {
		fmt.Println("test2")
		qb = qb.Where("pr.price <= ? ", strconv.Itoa(f.MaxPrice))
		sc = sc.Where("pr.price <= ? ", strconv.Itoa(f.MaxPrice))
	} else if len(strconv.Itoa(f.MinPrice)) != 1 && len(strconv.Itoa(f.MaxPrice)) != 1 {
		fmt.Println("test3")
		qb = qb.Where(squirrel.And{
			squirrel.GtOrEq{"pr.price": f.MinPrice},
			squirrel.LtOrEq{"pr.price": f.MaxPrice},
		})
		sc = sc.Where(squirrel.And{
			squirrel.GtOrEq{"pr.price": f.MinPrice},
			squirrel.LtOrEq{"pr.price": f.MaxPrice},
		})
	}

	if len(strconv.Itoa(f.MinStock)) != 1 && len(strconv.Itoa(f.MaxStock)) == 1 {
		qb = qb.Where("pr.price >= ? ", strconv.Itoa(f.MinStock))
		sc = sc.Where("pr.price >= ? ", strconv.Itoa(f.MinStock))
	} else if len(strconv.Itoa(f.MinStock)) == 1 && len(strconv.Itoa(f.MaxStock)) != 1 {
		qb = qb.Where("pr.price <= ? ", strconv.Itoa(f.MaxStock))
		sc = sc.Where("pr.price <= ? ", strconv.Itoa(f.MaxStock))
	} else if len(strconv.Itoa(f.MinStock)) != 1 && len(strconv.Itoa(f.MaxStock)) != 1 {
		qb = qb.Where(squirrel.And{
			squirrel.GtOrEq{"pr.price": f.MinStock},
			squirrel.LtOrEq{"pr.price": f.MaxStock},
		})
		sc = sc.Where(squirrel.And{
			squirrel.GtOrEq{"pr.price": f.MinStock},
			squirrel.LtOrEq{"pr.price": f.MaxStock},
		})
	}

	qb, pag := Paginate(sc, qb, *p)

	rows, err := qb.Query()
	if err != nil {
		log.Printf("Query rows failed: %v", err)
		return result, pag, errors.New("something wrong happened")
	}

	for rows.Next() {
		var r model.Book
		if err := rows.Scan(&r.ID, &r.UUID, &r.AuthorID, &r.Name, &r.Price, &r.Stock); err != nil {
			log.Printf("scan rows failed: %v", err)
			return result, pag, errors.New("something wrong happened")
		}

		result = append(result, r)
	}

	if len(p.Perpage) == 0 {
		defaultPerPage := "10"
		pag.Perpage = &defaultPerPage
	} else {
		pag.Perpage = &p.Perpage
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].ID > result[j].ID
	})

	var hmp bool
	if len(result) > 0 {
		more := sc.Where(squirrel.Lt{"id": result[len(result)-1].ID})
		hmp = hasMorePages(more)
	} else {
		hmp = false
	}

	pag.HasMorePages = &hmp
	if hmp {
		first := curs{result[len(result)-1].ID, *pag.HasMorePages}
		enc := encodeCursor(first)
		pag.NextCursor = &enc
	} else {
		pag.NextCursor = nil
	}

	if len(result) > 0 {
		sc = sc.Where(squirrel.Gt{"id": result[0].ID})
		if isFirstPage(sc) {
			pag.PrevCursor = nil
		} else {
			c := curs{result[0].ID, false}
			enc := encodeCursor(c)
			pag.PrevCursor = &enc
		}
	} else {
		pag.PrevCursor = nil
	}

	return result, pag, nil
}

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
	_, err := mysqlQB().Update("books").Set("author_id", bm.AuthorID).Set("name", bm.Name).Set("price", bm.Price).Where(squirrel.Eq{"uuid": uuid.UUID}).Exec()

	if err != nil {
		log.Printf("cannot update stock -> Error: %v", err)
		return errors.New("something wrong happened")
	} else {
		log.Printf("Success update Books")
	}

	return nil
}

func (br *bookRepository) DeleteBook(uuid dto.GetUUID) error {
	_, err := mysqlQB().Delete("books").Where(squirrel.Eq{"uuid": uuid.UUID}).Exec()
	if err != nil {
		log.Printf("cannot delete book -> Error: %v", err)
		return errors.New("something wrong happened")
	} else {
		log.Printf("Success delete Books")
	}

	return nil
}
