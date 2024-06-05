// Generated by gen.py. DO NOT modified it
package inject

import "github.com/gin-gonic/gin"

func Wrap4[T1, T2, T3, T4 any](f func(*gin.Context, T1, T2, T3, T4)) func(*gin.Context) {
	getter1, closer1 := processArg[T1]()
	getter2, closer2 := processArg[T2]()
	getter3, closer3 := processArg[T3]()
	getter4, closer4 := processArg[T4]()
	return func(c *gin.Context) {
		v1 := getter1()
		defer closer1(v1)
		v2 := getter2()
		defer closer2(v2)
		v3 := getter3()
		defer closer3(v3)
		v4 := getter4()
		defer closer4(v4)
		f(c, v1, v2, v3, v4)
	}
}
