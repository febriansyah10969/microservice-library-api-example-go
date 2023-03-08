package repository

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

type BookRepository interface {
	GetBooks(c *gin.Context, f *helper.Filter, p *helper.InPage) ([]model.Book, *helper.Pagination, error)
	CreateBook(bm model.Book) error
	UpdateBook(uuid dto.GetUUID, bm model.Book) error
}

type bookRepository struct {
	SQL  *sql.DB
	Sqlx *sqlx.DB
}
