package controllers

import (
	"errors"
	"net/http"

	"example.com/m/v2/models"
	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		token, err := services.LoginAdmin(c.Request.Context(), payload)
		if err != nil {
			if errors.Is(err, services.ErrInvalidCredentials) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not process login"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}