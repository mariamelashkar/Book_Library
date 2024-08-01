package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bookstore/internal/redis"

)
func BorrowBook(c *gin.Context) {
	username := c.MustGet("username").(string)
	role := c.MustGet("role").(string)
	if role != "user" {
		c.JSON(http.StatusForbidden, gin.H{"error": "only normal users can borrow books"})
		return
	}

	name := c.Param("Name")
	book, err := SearchForBook(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "The book is not found"})
		return
	}

	accessToken, err := redis.GenerateTokenWithExpiry(username, role, book.AccessTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "book": book})
}