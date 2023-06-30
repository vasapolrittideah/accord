package response

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type InvalidField struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func Fail(data []InvalidField) Response {
	return Response{
		Status: "fail",
		Data:   data,
	}
}

func Error(msg string) Response {
	return Response{
		Status:  "error",
		Message: msg,
	}
}

func Success(data interface{}) Response {
	return Response{
		Status: "success",
		Data:   data,
	}
}
