package errs

import "net/http"

type HttpError struct {
	error
	StatusCode int
}

func NewHttpError(statusCode int, err error) HttpError {
	return HttpError{
		error:      err,
		StatusCode: statusCode,
	}
}

func BadRequestError(err error) HttpError {
	return NewHttpError(http.StatusBadRequest, err)
}
