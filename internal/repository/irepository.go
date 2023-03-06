package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type BookRepository interface {
}

type bookRepository struct {
	SQL  *sql.DB
	Sqlx *sqlx.DB
}
