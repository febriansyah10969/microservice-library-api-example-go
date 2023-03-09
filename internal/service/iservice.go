package service

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
	"gitlab.com/p9359/backend-prob/febry-go/internal/repository"
)

type BookService interface {
	GetBooks(c *gin.Context, fillter *helper.Filter, paginate *helper.InPage) ([]model.Book, *helper.Pagination, error)
	CreateBook(c *gin.Context, rev dto.BookRequest) ([]string, error)
	UpdateBook(c *gin.Context, uuid dto.GetUUID, rev dto.BookRequest) ([]string, error)
	DeleteBook(uuid dto.GetUUID) error

	IncreaseStock(uri dto.GetUUID, req dto.StockRequest) error
	DecreaseStock(uri dto.GetUUID, req dto.StockRequest) error
}

type bookService struct {
	dao repository.DAO
}

func NewBookService(dao repository.DAO) BookService {
	return &bookService{dao}
}
