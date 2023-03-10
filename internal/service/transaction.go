package service

import (
	"github.com/google/uuid"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (bs *bookService) AddToCart(req dto.TransactionRequest, book model.Book, user model.User) error {
	dao := bs.dao.NewGeneralRepository()

	transaction := model.Transaction{}

	transaction.UUID = uuid.NewString()
	transaction.CodeTrx = generateInvoiceNumber()
	transaction.UserID = user.ID
	transaction.Days = req.Days
	transaction.Status = 1 // on cart
	transaction.FinalPrice = book.Price

	transaction_id, errCreateUserTransaction := dao.CreateUserTransaction(transaction)
	if errCreateUserTransaction != nil {
		return errCreateUserTransaction
	}

	bookTransaction := model.BookTransaction{}

	bookTransaction.TrxID = transaction_id
	bookTransaction.UUID = uuid.NewString()
	bookTransaction.BookID = book.ID
	bookTransaction.Qty = req.Qty
	bookTransaction.Price = book.Price // on cart

	dao.CreateBookTransaction(bookTransaction)

	book_uuid := dto.GetUUID{
		UUID: book.UUID,
	}

	book_qty := dto.StockRequest{
		Qty: req.Qty,
	}

	err := ControlStock(bs, book_uuid, book_qty, 5)
	if err != nil {
		return err
	}

	return nil
}
