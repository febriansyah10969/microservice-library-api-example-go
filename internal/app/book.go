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

func (ba *BookApp) GetListBook(c *gin.Context) {
	c.JSON(http.StatusOK, "testing")
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

	_, err := ba.BookService.CreateBook(c, req)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menambahkan buku", nil, req, nil)
	c.JSON(http.StatusOK, response)
}
