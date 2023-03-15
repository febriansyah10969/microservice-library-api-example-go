package service

import (
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

func (bs *bookService) GetBooks(fillter *helper.Filter, paginate *helper.InPage) ([]dto.BookResponse, *helper.Pagination, error) {
	data := []dto.BookResponse{}

	dao := bs.dao.NewGeneralRepository()
	result, pag, err := dao.GetBooks(fillter, paginate)
	if err != nil {
		return data, pag, err
	}

	for _, book := range result {
		response := dto.BookResponse{
			UUID:     book.UUID,
			AuthorID: book.AuthorID,
			Name:     book.Name,
			Price:    book.Price,
		}

		for _, category := range book.Category {
			catResponse := dto.BookCategoriesResponse{
				ID:   category.ID,
				Name: category.Name,
			}

			response.Category = append(response.Category, catResponse)
		}

		data = append(data, response)
	}

	return data, pag, nil
}

func (bs *bookService) CreateBook(rev dto.BookRequest) ([]string, error) {
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

func (bs *bookService) UpdateBook(uuid dto.GetUUID, rev dto.BookRequest) ([]string, error) {
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
