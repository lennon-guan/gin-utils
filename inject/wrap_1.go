// Generated by gen.py. DO NOT modified it
package inject

import "github.com/gin-gonic/gin"

func Wrap1[T1 any](f func(*gin.Context, T1)) func(*gin.Context) {
	getter1, closer1 := processArg[T1]()
	return func(c *gin.Context) {
		v1 := getter1(c)
		if closer1 != nil {
			defer closer1(v1)
		}
		f(c, v1)
	}
}
