package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

type BookRepository interface {
	GetUser(id int) (model.User, error)

	GetBookByID(id int) (model.Book, error)
	GetBook(book_uuid dto.GetUUID) (model.Book, error)
	GetBooks(f *helper.Filter, p *helper.InPage) ([]model.Book, *helper.Pagination, error)
	CreateBook(bm model.Book) error
	UpdateBook(uuid dto.GetUUID, bm model.Book) error
	DeleteBook(uuid dto.GetUUID) error

	CategoryDetail(getCategoryID dto.GetCategoryID) (model.Category, error)

	GetBookHistory(uuid dto.GetUUID, p *helper.InPage) ([]model.BookHistory, *helper.Pagination, error)
	GetCurrentStock(uuid dto.GetUUID) (model.Book, error)
	UpdateStock(uuid dto.GetUUID, stock int) error
	CreateBookHistory(mbh model.BookHistory) error

	GetTransaction(string) (model.Transaction, error)
	GetTransactions(f *helper.Filter, p *helper.InPage) ([]model.Transaction, *helper.Pagination, error)

	CreateUserTransaction(transaction model.Transaction) (int, error)
	CreateBookTransaction(transaction model.BookTransaction) error

	UpdateUserTransaction(trans_id int, transaction model.Transaction) error
	UpdateBookTransaction(trans_id int, transaction model.BookTransaction) error
}

type bookRepository struct {
	SQL  *sql.DB
	Sqlx *sqlx.DB
}
