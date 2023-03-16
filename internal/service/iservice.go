package service

import (
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
	"gitlab.com/p9359/backend-prob/febry-go/internal/repository"
)

type BookService interface {
	GetUser(id int) (model.User, error)

	GetBook(book_uuid dto.GetUUID) (model.Book, error)
	GetBookDetail(book_uuid dto.BookUUID) (dto.BookResponse, error)
	GetBooks(filter *helper.Filter, paginate *helper.InPage) ([]dto.BookResponse, *helper.Pagination, error)
	CreateBook(rev dto.BookRequest) ([]string, error)
	UpdateBook(uuid dto.GetUUID, rev dto.BookRequest) ([]string, error)
	DeleteBook(uuid dto.GetUUID) error

	GetTransactions(filter *helper.TrxFilter, paginate *helper.InPage) ([]dto.TransactionResponse, *helper.Pagination, error)

	CategoryDetail(categoryID dto.GetCategoryID) ([]dto.CategoryDetailResponse, error)

	GetBookHistory(uri dto.GetUUID, p *helper.InPage) ([]dto.BookHistoriesResponse, *helper.Pagination, error)
	IncreaseStock(uri dto.GetUUID, req dto.StockRequest) error
	DecreaseStock(uri dto.GetUUID, req dto.StockRequest) error

	AddToCart(req dto.TransactionRequest, book model.Book, user model.User) error
	OnBorrow(req dto.TransactionRequest, book model.Book, user model.User) error
	Finish(req dto.TransactionUUIDRequest) error
	Cancel(req dto.TransactionUUIDRequest) error
}

type bookService struct {
	dao repository.DAO
}

func NewBookService(dao repository.DAO) BookService {
	return &bookService{dao}
}
