package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ReviewRepository interface {
}

type reviewRepository struct {
	SQL  *sql.DB
	Sqlx *sqlx.DB
}
