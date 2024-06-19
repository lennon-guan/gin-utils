package inject

import (
	"fmt"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

type InjectFunc[T any] func(*gin.Context) T

// AddInjector registers a constructor function for type T in the dependency injection system.
// f: A constructor function that returns an instance of type T.
func AddInjector[T any](f InjectFunc[T]) {
	var v T
	injectTypeMap[reflect.TypeOf(v)] = f
}

var injectTypeMap = make(map[reflect.Type]any)

type closerWithoutReturnError interface {
	Close()
}

// Wrap wraps a handler function that takes *gin.Context and an argument of type T1.
// f: A handler function that takes *gin.Context and an argument of type T1.
func Wrap[T1 any](f func(*gin.Context, T1)) func(*gin.Context) {
	return Wrap1(f)
}

// processArg retrieves and generates an instance of type T, while also returning an appropriate closer function.
// Returns:
// - getter: A function that generates an instance of type T.
// - closer: A function used to close the generated instance.
func processArg[T any]() (getter InjectFunc[T], closer func(any)) {
	var v T
	vt := reflect.TypeOf(v)
	getter, is := injectTypeMap[vt].(InjectFunc[T])
	if !is {
		panic(fmt.Sprintf("inject type %s not registered", vt))
	}
	var vi any = v
	switch vi.(type) {
	case io.Closer:
		closer = closeIoCloser
	case closerWithoutReturnError:
		closer = closeCloserWithoutReturnError
	default:
		closer = nil
	}
	return
}

// closeIoCloser closes an instance that implements the io.Closer interface.
// v: The instance to be closed.
func closeIoCloser(v any) {
	v.(io.Closer).Close()
}

// closeCloserWithoutReturnError closes an instance that implements the closerWithoutReturnError interface.
// v: The instance to be closed.
func closeCloserWithoutReturnError(v any) {
	v.(closerWithoutReturnError).Close()
}
