package dto

type BookRequest struct {
	AuthorID int    `form:"author_id" json:"author_id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required,min=3"`
	Price    int    `form:"price" json:"price" binding:"required"`
}
