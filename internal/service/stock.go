package service

import (
	"errors"

	"github.com/google/uuid"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (bs *bookService) GetBookHistory(uri dto.GetUUID) (dto.BookHistoryResponse, error) {
	dao := bs.dao.NewGeneralRepository()

	data, err := dao.GetBookHistory(uri)

	if err != nil {
		return dto.BookHistoryResponse{}, err
	}

	getBook := dto.BookHistoryResponse{}
	getBookHistories := []dto.BookHistories{}

	for _, book := range data {
		getBook.UUID = book.UUID
		getBook.Name = book.Name
		getBook.Stock = book.Stock

		bookHistory := dto.BookHistories{
			UUID:   book.BookHistory.UUID,
			BookID: book.BookHistory.BookID,
			Qty:    book.BookHistory.Qty,
			Type:   book.BookHistory.Type,
		}

		getBookHistories = append(getBookHistories, bookHistory)

	}

	getBook.BookHistories = getBookHistories

	return getBook, nil
}

func (bs *bookService) IncreaseStock(uri dto.GetUUID, req dto.StockRequest) error {
	// Type 1 for increase
	err := ControlStock(bs, uri, req, 1)
	if err != nil {
		return err
	}

	return nil
}

func (bs *bookService) DecreaseStock(uri dto.GetUUID, req dto.StockRequest) error {
	// Type 2 for decrease
	err := ControlStock(bs, uri, req, 2)
	if err != nil {
		return err
	}

	return nil
}

func ControlStock(bs *bookService, uri dto.GetUUID, req dto.StockRequest, getType int) error {
	dao := bs.dao.NewGeneralRepository()
	data, errStock := dao.GetCurrentStock(uri)
	if errStock != nil {
		return errStock
	}

	stock := 0
	if getType == 1 {
		stock = data.Stock + req.Qty
	} else if getType == 2 {
		stock = data.Stock - req.Qty
		if stock <= 0 {
			return errors.New("Stok tidak cukup untuk dikurangi.")
		}
	}

	errIncrease := dao.UpdateStock(uri, stock)
	if errIncrease != nil {
		return errIncrease
	}

	bookHistory := model.BookHistory{}
	bookHistory.UUID = uuid.NewString()
	bookHistory.BookID = data.ID
	bookHistory.Qty = stock
	bookHistory.Type = 1

	errCreateBookHistory := dao.CreateBookHistory(bookHistory)
	if errCreateBookHistory != nil {
		return errCreateBookHistory
	}

	return nil
}
