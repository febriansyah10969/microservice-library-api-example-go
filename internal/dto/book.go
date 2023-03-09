package dto

type GetUUID struct {
	UUID string `uri:"uuid" binding:"required"`
}

type BookRequest struct {
	AuthorID int    `form:"author_id" json:"author_id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required,min=3"`
	Price    int    `form:"price" json:"price" binding:"required"`
}

type BookResponse struct {
	UUID     string `json:"uuid"`
	AuthorID int    `json:"author_id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
}
