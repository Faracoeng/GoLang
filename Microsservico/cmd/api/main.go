package main

import (
	"github.com/Faracoeng/GoLang/ms-categories/cmd/api/controllers"
	"github.com/Faracoeng/GoLang/ms-categories/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := repositories.InitDb()
	if err != nil {
		panic(err)
	}

	categoryRepo := repositories.NewCategoryRepository(db)

	// Injeta a interface no controller
	controllers.InitCategoryController(categoryRepo)

	router := gin.Default()
	CategoryRoutes(router)
	router.Run()
}
