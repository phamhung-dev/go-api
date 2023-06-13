package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{Data: data, Paging: nil, Filter: nil}
}

type errorResponse struct {
	Error string `json:"error"`
}

func SimpleErrorResponse(err error) *errorResponse {
	return &errorResponse{Error: err.Error()}
}
