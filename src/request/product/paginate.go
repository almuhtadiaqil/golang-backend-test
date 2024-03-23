package request

type ProductPaginationRequest struct {
	PageIndex int                    `form:"page_index" json:"page_index"`
	PageSize  int                    `form:"page_size" json:"page_size"`
	Query     string                 `form:"query" json:"query"`
	SortBy    string                 `form:"sort_by" json:"sort_by"`
	Filter    map[string]interface{} `form:"filter" json:"filter"`
}
