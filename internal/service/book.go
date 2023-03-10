package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (bs *bookService) GetBook(book_uuid dto.GetUUID) (model.Book, error) {
	dao := bs.dao.NewGeneralRepository()
	getBook, err := dao.GetBook(book_uuid)
	if err != nil {
		return model.Book{}, err
	}

	return getBook, nil
}

func (bs *bookService) GetBooks(c *gin.Context, fillter *helper.Filter, paginate *helper.InPage) ([]model.Book, *helper.Pagination, error) {
	dao := bs.dao.NewGeneralRepository()
	result, pag, err := dao.GetBooks(c, fillter, paginate)
	if err != nil {
		return result, pag, err
	}

	return result, pag, nil
}

func (bs *bookService) CreateBook(c *gin.Context, rev dto.BookRequest) ([]string, error) {
	repo := bs.dao.NewGeneralRepository()

	var bookData = new(model.Book)

	bookData.UUID = uuid.NewString()
	bookData.AuthorID = int(rev.AuthorID)
	bookData.Name = rev.Name
	bookData.Price = int(rev.Price)
	bookData.Stock = int(0)

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

func (bs *bookService) DeleteBook(uuid dto.GetUUID) error {

	repo := bs.dao.NewGeneralRepository()
	err := repo.DeleteBook(uuid)
	if err != nil {
		return err
	}

	return nil
}
