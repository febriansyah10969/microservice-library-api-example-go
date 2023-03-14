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

type BookTransform struct {
	UUID         string `json:"uuid"`
	AuthorID     int    `json:"author_id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Stock        int    `json:"stock"`
	ParentID     string `json:"parent_id"`
	SubID        string `json:"sub_id"`
	ChildSubID   string `json:"child_sub_id"`
	ParentName   string `json:"parent_name"`
	SubName      string `json:"sub_name"`
	ChildSubName string `json:"child_sub_name"`
}
