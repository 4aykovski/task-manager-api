package response

const (
	StatusOk             = "OK"
	StatusError          = "Error"
	UnauthorizedErrorMsg = "unauthorized"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func UnauthorizedError() Response {
	return Error(UnauthorizedErrorMsg)
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
