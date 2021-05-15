package middleware

import (
	"kokow/go/common"
	"kokow/go/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("Authorization")

		//validate token formate
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 404,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//验证通过 获取userid
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//验证用户是否存在在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 403,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}
		//用户存在将user信息写入上下文
		c.Set("user", user)
		c.Next()

	}

}
