package exceptions

import "net/http"

type HttpException struct {
	StatusCode int
	Message    string
	Details    interface{}
}

func (err *HttpException) Error() string {
	return err.Message
}

func NewHttpException(statusCode int, message string, detail interface{}) *HttpException {
	return &HttpException{
		StatusCode: statusCode,
		Message:    message,
		Details:    detail,
	}
}

func NewNotFoundException(message string, detail interface{}) *HttpException {
	return &HttpException{
		StatusCode: http.StatusNotFound,
		Message:    message,
		Details:    detail,
	}

}

func NewBadRequestException(message string, detail interface{}) *HttpException {
	return &HttpException{
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Details:    detail,
	}

}

func NewUnprocessableEntityException(message string, detail interface{}) *HttpException {
	return &HttpException{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
		Details:    detail,
	}

}

func NewForbiddenException(message string, detail interface{}) *HttpException {
	return &HttpException{
		StatusCode: http.StatusForbidden,
		Message:    message,
		Details:    detail,
	}

}
