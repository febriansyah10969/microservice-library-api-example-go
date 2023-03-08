package service

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/repository"
)

type BookService interface {
	CreateBook(c *gin.Context, rev dto.BookRequest) ([]string, error)
	UpdateBook(c *gin.Context, uuid dto.GetUUID, rev dto.BookRequest) ([]string, error)
}

type bookService struct {
	dao repository.DAO
}

func NewBookService(dao repository.DAO) BookService {
	return &bookService{dao}
}
