package handlers

type Response struct {
	StatusCode int64
	Message    string
}

func NewResponse(statusCode int64, message string) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
	}
}
