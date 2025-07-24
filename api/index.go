package handler

import (
	"net/http"

	db "example.com/m/v2/DB"
	"example.com/m/v2/controllers"
	"example.com/m/v2/middleware"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	db.InitDB()

	router = gin.Default()
	router.GET("/menu", controllers.GetMenuHandler())
	router.POST("/admin/login", controllers.LoginHandler())

	adminGroup := router.Group("/admin")
	adminGroup.Use(middleware.AuthMiddleware())
	{
		adminGroup.POST("/menu", controllers.CreateMenuItemHandler())
		adminGroup.PUT("/menu/:id", controllers.UpdateMenuItemHandler())
		adminGroup.DELETE("/menu/:id", controllers.DeleteMenuItemHandler())
		adminGroup.POST("/categories", controllers.CreateCategoryHandler())
		adminGroup.GET("/categories", controllers.GetAllCategoriesHandler())
		adminGroup.PUT("/categories/:id", controllers.UpdateCategoryHandler())
		adminGroup.DELETE("/categories/:id", controllers.DeleteCategoryHandler())
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
