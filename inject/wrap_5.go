// Generated by gen.py. DO NOT modified it
package inject

import "github.com/gin-gonic/gin"

func Wrap5[T1, T2, T3, T4, T5 any](f func(*gin.Context, T1, T2, T3, T4, T5)) func(*gin.Context) {
	getter1, closer1 := processArg[T1]()
	getter2, closer2 := processArg[T2]()
	getter3, closer3 := processArg[T3]()
	getter4, closer4 := processArg[T4]()
	getter5, closer5 := processArg[T5]()
	return func(c *gin.Context) {
		v1 := getter1(c)
		if closer1 != nil {
			defer closer1(v1)
		}
		v2 := getter2(c)
		if closer2 != nil {
			defer closer2(v2)
		}
		v3 := getter3(c)
		if closer3 != nil {
			defer closer3(v3)
		}
		v4 := getter4(c)
		if closer4 != nil {
			defer closer4(v4)
		}
		v5 := getter5(c)
		if closer5 != nil {
			defer closer5(v5)
		}
		f(c, v1, v2, v3, v4, v5)
	}
}
