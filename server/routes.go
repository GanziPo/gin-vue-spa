package main

import (
	"kokow/go/controller"
	"kokow/go/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	categoryRoutes := r.Group("/categories")
	cateController := controller.NewCategoryController()
	categoryRoutes.POST("", cateController.Create)
	categoryRoutes.PUT("/:id", cateController.Edit)
	categoryRoutes.GET("/:id", cateController.View)
	categoryRoutes.DELETE("/:id", cateController.Remove)
	// categoryRoutes.DELET
	return r
}
