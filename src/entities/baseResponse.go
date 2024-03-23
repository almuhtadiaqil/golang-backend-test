package entities

type BaseResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}
