package controllers

// Response Structure
type Response struct {
	Status bool        `json:"status"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"data"`
}

// Lists Response list structure
type Lists struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

// APIResponse is helper for giving response
func APIResponse(status bool, objects interface{}, msg string) (r *Response) {
	r = &Response{Status: status, Data: objects, Msg: msg}
	return
}
