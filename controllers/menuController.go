package controllers

import (
	"net/http"

	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func GetMenuHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fullMenu, err := services.GetFullMenu(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch menu data"})
			return
		}
		c.JSON(http.StatusOK, fullMenu)
	}
}