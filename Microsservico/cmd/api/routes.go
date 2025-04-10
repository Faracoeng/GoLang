package main

import (
	"github.com/Faracoeng/GoLang/ms-categories/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.Engine) {
	categoryRoutes := router.Group("/categories")

	categoryRoutes.POST("/", controllers.CreateCategory)
	categoryRoutes.GET("/", controllers.GetCategories)
	categoryRoutes.GET("/:id", controllers.GetCategory)
	categoryRoutes.PUT("/:id", controllers.UpdateCategory)
	categoryRoutes.DELETE("/:id", controllers.DeleteCategory)
}
