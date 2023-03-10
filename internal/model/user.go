package model

import "time"

// table name for book history
const USER string = "users"

type User struct {
	ID          int        `db:"id"`
	UUID        string     `db:"uuid"`
	Name        string     `db:"book_id"`
	Pin         string     `db:"qty"`
	Email       string     `db:"qty"`
	Phone       string     `db:"type"`
	DateOfBirth *time.Time `db:"type"`
}
