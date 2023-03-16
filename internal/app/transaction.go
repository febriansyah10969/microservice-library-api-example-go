package app

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/middleware"
)

func (ba *BookApp) GetTransactions(c *gin.Context) {
	fillter := new(helper.Filter)

	if err := c.BindQuery(&fillter); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menampilkan transaksi", nil, []dto.BookResponse{}, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inPag := new(helper.InPage)
	if err := c.ShouldBindWith(&inPag, binding.Query); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menampilkan transaksi", nil, []dto.BookResponse{}, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, pag, errGetTransactions := ba.BookService.GetTransactions(fillter, inPag)
	if errGetTransactions != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menampilkan transaksi", nil, nil, errGetTransactions.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menampilkan transaksi", pag, transactions, nil)
	c.JSON(http.StatusOK, response)
}

func (ba *BookApp) AddToCart(c *gin.Context) {
	payload := c.MustGet("payload").(middleware.Auth)
	getUser, errUserService := ba.BookService.GetUser(payload.ID)
	if errUserService != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, errUserService.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	req := dto.TransactionRequest{}
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		fmt.Println(&req) // log the error
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, nil)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				var err helper.FieldValidation

				err.Attribute = strings.ToLower(fe.Field())
				err.Text = fmt.Sprintf("%v harus diisi", fe.Field())

				response.Error = append(response.Error, err)
			}
		}

		c.JSON(http.StatusOK, response)
		return
	}

	uri := dto.GetUUID{
		UUID: req.BookUUID,
	}

	getBook, errBook := ba.BookService.GetBook(uri)
	if errBook != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, errBook.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errBookService := ba.BookService.AddToCart(req, getBook, getUser)
	if errBookService != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, errBookService.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menambahkan buku ke keranjang", nil, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (ba *BookApp) OnBorrow(c *gin.Context) {
	payload := c.MustGet("payload").(middleware.Auth)
	getUser, errUserService := ba.BookService.GetUser(payload.ID)
	if errUserService != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, errUserService.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	req := dto.TransactionRequest{}
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		fmt.Println(&req) // log the error
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, nil)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				var err helper.FieldValidation

				err.Attribute = strings.ToLower(fe.Field())
				err.Text = fmt.Sprintf("%v harus diisi", fe.Field())

				response.Error = append(response.Error, err)
			}
		}

		c.JSON(http.StatusOK, response)
		return
	}

	uri := dto.GetUUID{
		UUID: req.BookUUID,
	}

	getBook, errBook := ba.BookService.GetBook(uri)
	if errBook != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, errBook.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	errBookService := ba.BookService.OnBorrow(req, getBook, getUser)
	if errBookService != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, errBookService.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menambahkan buku ke keranjang", nil, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (ba *BookApp) Finish(c *gin.Context) {
	req := dto.TransactionUUIDRequest{}
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		fmt.Println(&req) // log the error
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, nil)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				var err helper.FieldValidation

				err.Attribute = strings.ToLower(fe.Field())
				err.Text = fmt.Sprintf("%v harus diisi", fe.Field())

				response.Error = append(response.Error, err)
			}
		}

		c.JSON(http.StatusOK, response)
		return
	}

	errBookService := ba.BookService.Finish(req)
	if errBookService != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku ke keranjang", nil, nil, errBookService.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menambahkan buku ke keranjang", nil, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (ba *BookApp) Cancel(c *gin.Context) {
	req := dto.TransactionUUIDRequest{}
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		fmt.Println(&req) // log the error
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal mengembalikan buku", nil, nil, nil)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				var err helper.FieldValidation

				err.Attribute = strings.ToLower(fe.Field())
				err.Text = fmt.Sprintf("%v harus diisi", fe.Field())

				response.Error = append(response.Error, err)
			}
		}

		c.JSON(http.StatusOK, response)
		return
	}

	errBookService := ba.BookService.Cancel(req)
	if errBookService != nil {

		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal mengembalikan buku", nil, nil, errBookService.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil mengembalikan buku", nil, nil, nil)
	c.JSON(http.StatusOK, response)
}
