package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bookstore/models"

)
func CreateBook(c *gin.Context) {
	role := c.MustGet("role").(string)
	if role != "author" {
		c.JSON(http.StatusForbidden, gin.H{"error": "only authors can create books"})
		return
	}

	var newBook models.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newBook.AccessTime <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid access time"})
		return
	}

	models.Books = append(models.Books, newBook)
	c.JSON(http.StatusCreated, newBook)
}
