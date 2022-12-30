package errs

import (
	"errors"
	"net/http"
	"zamannow/go-rest-api/dto/responses"

	"github.com/gin-gonic/gin"
)

func Handler(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	causeErr := getErrorCause(err)

	var code int
	switch e := causeErr.(type) {
	case *defaultError:
		code = e.code
	default:
		// returnedErr = NewServerError("internal server error")
		code = http.StatusInternalServerError
	}

	genResp(ctx, code, causeErr)
}

func genResp(ctx *gin.Context, code int, err error) {
	resp := responses.ErrorResponse{
		Status: code,
		Error:  err,
	}
	ctx.AbortWithStatusJSON(code, resp)
}

func getErrorCause(err error) error {
	oldestErr := err
	for {
		tempErr := errors.Unwrap(oldestErr)
		if tempErr == nil {
			break
		}

		oldestErr = tempErr
	}

	return oldestErr
}
