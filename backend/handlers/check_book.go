package handlers
import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func Checkoutbookbyid(c *gin.Context) {
	name := c.Param("Name")
	book, err := SearchForBook(name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The book is Not found"})
		return
	}
	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}
func Addbookbyid(c *gin.Context) {
	name := c.Param("Name")
	book, err := SearchForBook(name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "The book is Not found"})
		return
	}
	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}