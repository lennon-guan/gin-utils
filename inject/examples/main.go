package main

import (
	"fmt"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/lennon-guan/gin-utils/inject"
)

func main() {
	var counter atomic.Int64

	// 通过AddInjector注入int64类型的注入方法
	// Add an injector method for int64 type using AddInjector
	inject.AddInjector(func(*gin.Context) int64 {
		return counter.Add(1)
	})

	// 通过AddInjector注入SomeResource类型的注入方法
	// SomeResource类型实现了io.Closer接口，因此在HandlerFunc返回的时候，会自动执行其Close方法
	// Add an injector method for SomeResource type using AddInjector
	// SomeResource implements the io.Closer interface, so its Close method will be automatically called when the HandlerFunc returns
	inject.AddInjector(func(*gin.Context) SomeResource {
		fmt.Println("made a new SomeResource!")
		return SomeResource{}
	})

	engine := gin.Default()

	// 这里通过inject.Wrap2方法，包装了一个有两个待注入参数的HandlerFunc
	// 支持Wrap1 - Wrap9，inject.Wrap相当于inject.Wrap1
	// Here, a HandlerFunc with two injectable parameters is wrapped using the inject.Wrap2 method
	// Supports Wrap1 - Wrap9, inject.Wrap is equivalent to inject.Wrap1
	engine.GET("/", inject.Wrap2(func(c *gin.Context, counter int64, res SomeResource) {
		c.String(200, "req %d", counter)
	}))
	engine.Run(":28080")
}

type SomeResource struct{}

func (SomeResource) Close() error {
	fmt.Println("SomeResource as io.Closer released!")
	return nil
}
