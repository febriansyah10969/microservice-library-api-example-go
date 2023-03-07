package model

// table name for book
const BOOK string = "books"

type Book struct {
	UUID     string `db:"uuid"`
	AuthorID int    `db:"author_id"`
	Name     string `db:"name"`
	Price    int    `db:"price"`
	Stock    int    `db:"stock"`
}
