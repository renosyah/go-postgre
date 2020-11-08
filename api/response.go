package api

type (
	Response struct {
		Status int         `json:"status"`
		Errors []Error     `json:"errors"`
		Data   interface{} `json:"data"`
	}
)
