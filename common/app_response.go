package common

type SuccessRes struct {
	Data   interface{} `json:"data,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
	Paging interface{} `json:"paging,omitempty"`
}

func NewSuccessResponse(data, filter, paging interface{}) *SuccessRes {
	return &SuccessRes{
		Data:   data,
		Filter: filter,
		Paging: paging,
	}
}
func SimpleSuccessResponse(data interface{}) *SuccessRes {
	return &SuccessRes{
		Data:   data,
		Filter: nil,
		Paging: nil,
	}
}
