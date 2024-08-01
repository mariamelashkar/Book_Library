package handlers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"bookstore/models"
	"golang.org/x/crypto/bcrypt"
	"bookstore/internal/redis"
)
func LoginUser(c *gin.Context) {
	var userCredentials models.User
	if err := c.BindJSON(&userCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, user := range models.Users {
		if user.Username == userCredentials.Username {
			if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userCredentials.Password)); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
				return
			}

			tokenString, err := redis.GenerateToken(user.Username, user.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": tokenString})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
}