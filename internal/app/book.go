package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ba *BookApp) GetListBook(c *gin.Context) {
	c.JSON(http.StatusOK, "testing")
}
