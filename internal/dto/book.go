package dto

type GetUUID struct {
	UUID string `uri:"uuid" binding:"required"`
}

type BookRequest struct {
	AuthorID int    `form:"author_id" json:"author_id" binding:"required"`
	Name     string `form:"name" json:"name" binding:"required,min=3"`
	Price    int    `form:"price" json:"price" binding:"required"`
	Stock    int    `form:"stock" json:"stock" binding:"required"`
}
