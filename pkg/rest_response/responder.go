package rest_response

type responder struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *responder {
	return &responder{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *responder {
	return NewSuccessResponse(data, nil, nil)
}
