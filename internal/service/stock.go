package service

import (
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
)

func (bs *bookService) GetBookHistory(uri dto.GetUUID, p *helper.InPage) ([]dto.BookHistoriesResponse, *helper.Pagination, error) {
	dao := bs.dao.NewGeneralRepository()

	data, pag, err := dao.GetBookHistory(uri, p)

	if err != nil {
		return []dto.BookHistoriesResponse{}, pag, err
	}

	getBookHistories := []dto.BookHistoriesResponse{}

	for _, book := range data {
		bookHistory := dto.BookHistoriesResponse{
			UUID:   book.UUID,
			BookID: book.BookID,
			Qty:    book.Qty,
			Type:   book.Type,
		}

		getBookHistories = append(getBookHistories, bookHistory)

	}

	return getBookHistories, pag, nil
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
