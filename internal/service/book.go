package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (bs *bookService) CreateBook(c *gin.Context, rev dto.BookRequest) ([]string, error) {
	repo := bs.dao.NewGeneralRepository()

	var bookData = new(model.Book)

	bookData.UUID = uuid.NewString()
	bookData.AuthorID = int(rev.AuthorID)
	bookData.Name = rev.Name
	bookData.Price = int(rev.Price)
	bookData.Stock = int(rev.Stock)

	err := repo.CreateBook(*bookData)
	if err != nil {
		return []string{}, err
	}

	return []string{}, nil
}

func (bs *bookService) UpdateBook(c *gin.Context, uuid dto.GetUUID, rev dto.BookRequest) ([]string, error) {
	// var uri dto.GetUUID

	repo := bs.dao.NewGeneralRepository()

	var bookData = new(model.Book)

	bookData.AuthorID = int(rev.AuthorID)
	bookData.Name = rev.Name
	bookData.Price = int(rev.Price)

	err := repo.UpdateBook(uuid, *bookData)
	if err != nil {
		return []string{}, err
	}

	return []string{}, nil
}
