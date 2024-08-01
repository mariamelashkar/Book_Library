package handlers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bookstore/models"
	"golang.org/x/crypto/bcrypt"

)
func RegisterUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUser.Password = string(hashedPassword)
	models.Users = append(models.Users, newUser)
	c.JSON(http.StatusCreated, newUser)
}