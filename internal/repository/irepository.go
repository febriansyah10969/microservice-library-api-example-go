package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

type BookRepository interface {
	CreateBook(bm model.Book) error
	UpdateBook(uuid dto.GetUUID, bm model.Book) error
}

type bookRepository struct {
	SQL  *sql.DB
	Sqlx *sqlx.DB
}
