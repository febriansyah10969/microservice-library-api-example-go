package service

import (
	"github.com/google/uuid"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (bs *bookService) IncreaseStock(uri dto.GetUUID, req dto.StockRequest) error {
	dao := bs.dao.NewGeneralRepository()
	data, errStock := dao.GetCurrentStock(uri)
	if errStock != nil {
		return errStock
	}

	stock := data.Stock + req.Qty

	errIncrease := dao.IncreaseStock(uri, stock)
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
