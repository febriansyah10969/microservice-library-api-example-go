package service

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func generateInvoiceNumber() string {
	// Mengatur seed generator acak dengan waktu saat ini
	rand.Seed(time.Now().UnixNano())
	// Menghasilkan bilangan bulat acak antara 0 hingga 9
	randomInt := rand.Intn(8999) + 1000
	// Mendapatkan waktu saat ini dalam format YmdHis dan microsecond
	now := time.Now()
	str := now.Format("20060102150405")

	return "TRXID-" + str + strconv.Itoa(randomInt)
}

func ControlStock(bs *bookService, book_uuid dto.GetUUID, req dto.StockRequest, getType int) error {
	dao := bs.dao.NewGeneralRepository()
	data, errStock := dao.GetCurrentStock(book_uuid)
	if errStock != nil {
		return errStock
	}

	stock := 0
	if getType == 1 || getType == 4 {
		stock = data.Stock + req.Qty
	} else if getType == 2 || getType == 5 || getType == 6 {
		stock = data.Stock - req.Qty
		if stock <= 0 {
			return errors.New("Stok tidak mencukupi.")
		}
	}

	errIncrease := dao.UpdateStock(book_uuid, stock)
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
