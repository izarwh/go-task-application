package common

type CommonResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewCommonResponse(success bool, message string, data any) *CommonResponse {
	return &CommonResponse{
		Success: success,
		Message: message,
		Data:    data,
	}
}
