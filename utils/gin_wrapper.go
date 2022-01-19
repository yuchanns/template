package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BuildGinHandler 统一构造 gin 的路由函数
func BuildGinHandler(fn func(ctx *gin.Context) (interface{}, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, err := fn(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"msg":  err.Error(),
				"code": 500,
				"data": nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"code": 0,
			"data": data,
		})
	}
}
