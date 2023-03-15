package dto

type GetCategoryID struct {
	CategoryID string `uri:"id" binding:"required"`
}

type CategoryDetailResponse struct {
	ID       int    `json:"id"`
	ParentID *int   `json:"parent_id"`
	Name     string `json:"name"`
}
