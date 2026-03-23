package result

import (
	"errors"
	"fmt"
	"net/http"
)

type Void struct{}

type Result[T any] struct {
	error     error
	isSuccess bool
	code      int
	value     T
}

func (r Result[T]) IsSuccess() bool {
	return r.isSuccess
}

func (r Result[T]) Value() T {
	if !r.isSuccess {
		var zeroValue T
		return zeroValue
	}
	return r.value
}

func (r Result[T]) Code() int {
	return r.code
}

func (r Result[T]) Error() error {
	return r.error
}

func (r Result[T]) String() string {
	return fmt.Sprintf("Result[IsSuccess=%v, Code=%d, Error=%v, Value=%v]",
		r.isSuccess, r.code, r.error, r.value)
}

func OK[T any](value T) Result[T] {
	return Result[T]{
		error:     nil,
		isSuccess: true,
		code:      http.StatusOK,
		value:     value,
	}
}

func SimpleOk() Result[Void] {
	return OK(Void{})
}

func Fail[T any](err error) Result[T] {
	return Result[T]{
		error:     err,
		isSuccess: false,
		code:      http.StatusBadRequest,
	}
}

func SimpleFail(err error) Result[Void] {
	return Fail[Void](err)
}

func NotFound[T any](err error) Result[T] {
	return Result[T]{
		error:     err,
		isSuccess: false,
		code:      http.StatusNotFound,
	}
}

func SimpleNotFound() Result[Void] {
	return NotFound[Void](errors.New("Not found"))
}
