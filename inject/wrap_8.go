// Generated by gen.py. DO NOT modified it
package inject

import "github.com/gin-gonic/gin"

func Wrap8[T1, T2, T3, T4, T5, T6, T7, T8 any](f func(*gin.Context, T1, T2, T3, T4, T5, T6, T7, T8)) func(*gin.Context) {
	getter1, closer1 := processArg[T1]()
	getter2, closer2 := processArg[T2]()
	getter3, closer3 := processArg[T3]()
	getter4, closer4 := processArg[T4]()
	getter5, closer5 := processArg[T5]()
	getter6, closer6 := processArg[T6]()
	getter7, closer7 := processArg[T7]()
	getter8, closer8 := processArg[T8]()
	return func(c *gin.Context) {
		v1 := getter1()
		if closer1 != nil {
			defer closer1(v1)
		}
		v2 := getter2()
		if closer2 != nil {
			defer closer2(v2)
		}
		v3 := getter3()
		if closer3 != nil {
			defer closer3(v3)
		}
		v4 := getter4()
		if closer4 != nil {
			defer closer4(v4)
		}
		v5 := getter5()
		if closer5 != nil {
			defer closer5(v5)
		}
		v6 := getter6()
		if closer6 != nil {
			defer closer6(v6)
		}
		v7 := getter7()
		if closer7 != nil {
			defer closer7(v7)
		}
		v8 := getter8()
		if closer8 != nil {
			defer closer8(v8)
		}
		f(c, v1, v2, v3, v4, v5, v6, v7, v8)
	}
}
