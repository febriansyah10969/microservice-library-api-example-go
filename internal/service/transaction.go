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

func (bs *bookService) OnBorrow(req dto.TransactionRequest, book model.Book, user model.User) error {
	dao := bs.dao.NewGeneralRepository()

	// status 1 on cart
	// status 2 on borrow
	// status 3 finish
	// status 4 cancel

	// type in = 1
	// type out = 2
	// type finish = 3
	// type cancel = 4
	// type on cart = 5
	// type on borrow = 6

	// * Define the struct
	book_uuid := dto.GetUUID{UUID: book.UUID}
	book_qty := dto.StockRequest{Qty: req.Qty}

	if req.TransID == 0 {
		transaction := model.Transaction{}

		transaction.UUID = uuid.NewString()
		transaction.CodeTrx = generateInvoiceNumber()
		transaction.UserID = user.ID
		transaction.Days = req.Days
		transaction.Status = 2 // on Borrow
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
		bookTransaction.Price = book.Price // on Borrow

		dao.CreateBookTransaction(bookTransaction)
	} else {
		transaction := model.Transaction{}

		transaction.Days = req.Days
		transaction.Status = 2 // on Borrow

		errCreateUserTransaction := dao.UpdateUserTransaction(req.TransID, transaction)
		if errCreateUserTransaction != nil {
			return errCreateUserTransaction
		}

		bookTransaction := model.BookTransaction{}
		bookTransaction.Qty = req.Qty

		dao.UpdateBookTransaction(req.TransID, bookTransaction)

		errIncreaseStock := ControlStock(bs, book_uuid, book_qty, 4)
		if errIncreaseStock != nil {
			return errIncreaseStock
		}
	}

	errDecreaseStock := ControlStock(bs, book_uuid, book_qty, 6)
	if errDecreaseStock != nil {
		return errDecreaseStock
	}

	return nil
}
