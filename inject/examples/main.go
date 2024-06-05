package main

import (
	"fmt"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/lennon-guan/gin-utils/inject"
)

func main() {
	var counter atomic.Int64
	inject.AddInjector(func() int64 {
		return counter.Add(1)
	})
	inject.AddInjector(func() SomeResource {
		fmt.Println("made a new SomeResource!")
		return SomeResource{}
	})

	engine := gin.Default()
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
