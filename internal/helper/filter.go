package helper

type Filter struct {
	BookUUID string `form:"book_uuid"`
	BookID   int    `form:"book_id"`
	AuthorID int    `form:"author_id"`
	Name     string `form:"name"`
	MinPrice int    `form:"min_price"`
	MaxPrice int    `form:"max_price"`
	MinStock int    `form:"min_stock"`
	MaxStock int    `form:"max_stock"`
}

type TrxFilter struct {
	TrxID string `form:"trx_id"`
}

type Timezone struct {
	CurrentTime string `form:"makasi-timenow"`
	GMT         int    `form:"makasi-timezone"`
}
