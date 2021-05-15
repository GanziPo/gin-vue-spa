package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}
func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 403, data, msg)
}

// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 401, "msg": "手机号码不正确"})
// 		return
