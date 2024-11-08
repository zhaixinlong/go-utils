package util

import (
	"fmt"

	errors "github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/status"
)

const ERR_CODE_UNDEFINED = -1

func ErrCode(err error) int {
	type errCoder interface {
		ErrCode() int
	}
	code := ERR_CODE_UNDEFINED
	if err != nil {
		errCode, ok := err.(errCoder)
		if ok {
			code = errCode.ErrCode()
		}
	}
	return code
}

func WithErrCode(err error, code int) error {
	if err == nil {
		return nil
	}
	return &withCode{
		cause: err,
		code:  code,
	}
}

type withCode struct {
	cause error
	code  int
}

func (w *withCode) Error() string {
	return fmt.Sprintf("<%d>%s", w.code, w.cause)
}

func (w *withCode) ErrCode() int {
	return w.code
}

func WithKratosErrorCode(err error) (int32, string) {
	code := int32(-1)
	message := "An unknown error"
	e, ok := err.(*errors.Error)
	if ok {
		code = e.Code
		message = e.Message
	}
	return code, message
}

func WithGrpcErrorCode(err error) (int32, string) {
	code := int32(-1)
	message := "An unknown error"
	e, ok := status.FromError(err)
	if ok {
		code = int32(e.Code())
		message = e.Message()
	}
	return code, message
}
