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
)

// func (ba *BookApp) GetListBook(c *gin.Context) {
// 	c.JSON(http.StatusOK, "testing")
// }

func (ba *BookApp) GetListBook(c *gin.Context) {
	fillter := new(helper.Filter)

	if err := c.BindQuery(&fillter); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menampilkan daftar ulasan", nil, []dto.BookResponse{}, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	inPag := new(helper.InPage)
	if err := c.ShouldBindWith(&inPag, binding.Query); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menampilkan daftar ulasan", nil, []dto.BookResponse{}, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result, pag, err := ba.BookService.GetBooks(fillter, inPag)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menampilkan daftar Buku", nil, []dto.BookResponse{}, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menampilkan daftar Buku", pag, result, nil)
	c.JSON(http.StatusOK, response)
}

func (ba *BookApp) CreateBook(c *gin.Context) {
	req := dto.BookRequest{}

	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		fmt.Println(err) // log the error
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal Menambahkan buku", nil, nil, nil)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				fmt.Println(fe.Param())
				var err helper.FieldValidation
				if fe.Tag() == "min" {
					err.Attribute = strings.ToLower(fe.Field())
					err.Text = fmt.Sprintf("%v minimal %v karakter", fe.Field(), fe.Param())
				} else {
					err.Attribute = strings.ToLower(fe.Field())
					err.Text = fmt.Sprintf("%v harus diisi", fe.Field())
				}

				response.Error = append(response.Error, err)
			}
		}

		c.JSON(http.StatusOK, response)
		return
	}

	_, err := ba.BookService.CreateBook(req)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menambahkan buku", nil, nil, nil)
	c.JSON(http.StatusOK, response)
}

func (ba *BookApp) UpdateBook(c *gin.Context) {
	uri := dto.GetUUID{}
	if err := c.ShouldBindUri(&uri); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal Menambah buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	req := dto.BookRequest{}
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		fmt.Println(err) // log the error
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal Menambahkan buku", nil, nil, nil)

		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				fmt.Println(fe.Param())
				var err helper.FieldValidation
				if fe.Tag() == "min" {
					err.Attribute = strings.ToLower(fe.Field())
					err.Text = fmt.Sprintf("%v minimal %v karakter", fe.Field(), fe.Param())
				} else {
					err.Attribute = strings.ToLower(fe.Field())
					err.Text = fmt.Sprintf("%v harus diisi", fe.Field())
				}

				response.Error = append(response.Error, err)
			}
		}

		c.JSON(http.StatusOK, response)
		return
	}

	_, err := ba.BookService.UpdateBook(uri, req)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menambahkan buku", nil, req, nil)
	c.JSON(http.StatusOK, response)
}

func (ba *BookApp) DeleteBook(c *gin.Context) {
	uri := dto.GetUUID{}
	if err := c.ShouldBindUri(&uri); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menghapus buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err := ba.BookService.DeleteBook(uri)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menghapus buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menghapus buku", nil, nil, nil)
	c.JSON(http.StatusOK, response)
}
