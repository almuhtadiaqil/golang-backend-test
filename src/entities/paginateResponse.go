package entities

type PaginateResponse struct {
	BaseResponse
	PageIndex int   `json:"page_index"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
}
