package controller

import "github.com/gin-gonic/gin"

type RestController interface {
	Create(ctx *gin.Context)
	Edit(ctx *gin.Context)
	Remove(ctx *gin.Context)
	View(ctx *gin.Context)
}
