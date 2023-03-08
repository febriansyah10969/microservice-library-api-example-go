package helper

type Pagination struct {
	NextCursor   *string `json:"next_cursor"`
	PrevCursor   *string `json:"prev_cursor"`
	Perpage      *string `json:"perpage"`
	HasMorePages *bool   `json:"has_more_pages"`
}

type InPage struct {
	Perpage string  `form:"perpage"`
	Cursor  *string `form:"cursor"`
}
