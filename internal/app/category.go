package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
)

func (ba *BookApp) GetCategoryDetail(c *gin.Context) {
	categoryID := dto.GetCategoryID{}
	if err := c.ShouldBindUri(&categoryID); err != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal mendapatkan rincian kategori", nil, nil, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data, errCategoryDetail := ba.BookService.CategoryDetail(categoryID)
	if errCategoryDetail != nil {
		response := helper.APIResponse(http.StatusBadRequest, false, "Gagal mendapatkan rincian kategori", nil, data, errCategoryDetail.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(http.StatusOK, true, "Berhasil mendapatkan rincian kategori", nil, data, nil)
	c.JSON(http.StatusOK, response)
}
