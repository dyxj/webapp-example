package app

import (
	"fmt"
	"net/http"
)

type errorApp struct {
	Error   error
	Message string
	Code    int
}

func genAppError(err error, format string, v ...interface{}) *errorApp {
	return &errorApp{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    http.StatusInternalServerError,
	}
}
