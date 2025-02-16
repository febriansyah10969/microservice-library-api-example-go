package model

// table name for book
const BOOK string = "books"

type Book struct {
	ID          int    `db:"id"`
	UUID        string `db:"uuid"`
	AuthorID    int    `db:"author_id"`
	Name        string `db:"name"`
	Price       int    `db:"price"`
	Stock       int    `db:"stock"`
	BookHistory BookHistory
	Category    []Category
}

type Category struct {
	ID         *int    `db:"id"`
	CategoryID *int    `db:"category_id"`
	Name       *string `db:"name"`
	Category   *Category
}
