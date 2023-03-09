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

func (ba *BookApp) IncreaseBookStock(c *gin.Context) {
	uri := dto.GetUUID{}
	if err := c.ShouldBindUri(&uri); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal Menambah stok buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	req := dto.StockRequest{}
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		fmt.Println(err) // log the error
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal Menambahkan stok buku", nil, nil, nil)

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

	err := ba.BookService.IncreaseStock(uri, req)
	if err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal menambahkan stok buku", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil menambahkan stok buku", nil, req, nil)
	c.JSON(http.StatusOK, response)
}
