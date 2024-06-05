package inject

import (
	"fmt"
	"io"
	"reflect"

	"github.com/gin-gonic/gin"
)

func AddInjector[T any](f func() T) {
	var v T
	injectTypeMap[reflect.TypeOf(v)] = f
}

var injectTypeMap = make(map[reflect.Type]any)

type closerWithoutReturnError interface {
	Close()
}

func Wrap[T1 any](f func(*gin.Context, T1)) func(*gin.Context) {
	return Wrap1(f)
}

func processArg[T any]() (getter func() T, closer func(any)) {
	var v T
	vt := reflect.TypeOf(v)
	getter, is := injectTypeMap[vt].(func() T)
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
		closer = closeNoop
	}
	return
}

func closeIoCloser(v any) {
	v.(io.Closer).Close()
}

func closeCloserWithoutReturnError(v any) {
	v.(closerWithoutReturnError).Close()
}

func closeNoop(any) {}
