package service

import (
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
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
