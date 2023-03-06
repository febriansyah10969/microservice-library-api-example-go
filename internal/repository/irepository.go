package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type GeneralRepository interface {
}

type generalRepository struct {
	SQL  *sql.DB
	Sqlx *sqlx.DB
}
