package helper

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}
