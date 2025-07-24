package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"example.com/m/v2/models"
	"example.com/m/v2/repository"
	"example.com/m/v2/services"
	"github.com/gin-gonic/gin"
)

func CreateMenuItemHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.AdminMenuItemPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		createdItem, err := services.CreateMenuItem(c.Request.Context(), &payload)
		if err != nil {
			if errors.Is(err, services.ErrCategoryNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create menu item"})
			return
		}

		c.JSON(http.StatusCreated, createdItem)
	}
}

func DeleteMenuItemHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item ID"})
			return
		}

		err = repository.DeleteMenuItem(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete menu item"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "item deleted successfully"})
	}
}

func UpdateMenuItemHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item ID"})
			return
		}

		var payload models.AdminMenuItemPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}

		err = services.UpdateMenuItem(c.Request.Context(), id, &payload)
		if err != nil {
			if errors.Is(err, services.ErrMenuItemNotFound) || errors.Is(err, services.ErrCategoryNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update menu item"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "item updated successfully"})
	}
}

func CreateCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload models.AdminCategoryPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		newCategory, err := services.CreateCategory(c.Request.Context(), &payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create category"})
			return
		}
		c.JSON(http.StatusCreated, newCategory)
	}
}

func GetAllCategoriesHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		categories, err := services.GetAllCategories(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve categories"})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}

func UpdateCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		var payload models.AdminCategoryPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
			return
		}
		err := services.UpdateCategory(c.Request.Context(), id, &payload)
		if err != nil {
			if errors.Is(err, services.ErrCategoryNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update category"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "category updated successfully"})
	}
}

func DeleteCategoryHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
		err := services.DeleteCategory(c.Request.Context(), id)
		if err != nil {
			if errors.Is(err, services.ErrCategoryNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}
			if errors.Is(err, services.ErrCategoryInUse) {
				c.JSON(http.StatusConflict, gin.H{"error": err.Error()}) 
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete category"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "category deleted successfully"})
	}
}
