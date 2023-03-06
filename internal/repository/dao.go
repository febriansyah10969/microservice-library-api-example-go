package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DAO interface {
	NewGeneralRepository() BookRepository
}

type dao struct {
	sql  *sql.DB
	sqlx *sqlx.DB
}

var DB *sqlx.DB

func NewDAO(sql *sql.DB, sqlx *sqlx.DB) DAO {
	return &dao{sql, sqlx}
}

func mysqlQB() squirrel.StatementBuilderType {
	return squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question).RunWith(DB)
}

func NewGorm() (*gorm.DB, error) {
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", db_username, db_password, db_host, db_port, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("cannot intiation gorm, error: %v", err)
		return nil, err
	}

	return db, nil
}

func NewSQLDB() (*sqlx.DB, error) {
	var err error

	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_DATABASE")
	db_connection := os.Getenv("DB_CONNECTION")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local", db_username, db_password, db_host, db_port, db_name)
	DB, err = sqlx.Open(db_connection, dsn)
	if err != nil {
		log.Printf("err: %v", err)
		return nil, err
	}

	return DB, nil
}

func (d *dao) NewGeneralRepository() BookRepository {
	return &bookRepository{d.sql, d.sqlx}
}
