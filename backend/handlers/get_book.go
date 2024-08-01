package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bookstore/models"

)
func Getbook(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Books)

}

func Getbookbyname(c *gin.Context) {
	name := c.Param("Name") 
	book, err := SearchForBook(name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The book is Not found"}) 
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}