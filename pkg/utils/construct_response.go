package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func NewSuccessResponse(message string, data any) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func NewFailedResponse(message string, err error, data any) Response {
	return Response{
		Status:  false,
		Message: message,
		Error:   err.Error(),
		Data:    data,
	}
}
