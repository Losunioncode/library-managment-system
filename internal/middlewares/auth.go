package middlewares

import (
	"github.com/gin-gonic/gin"
	"github/losunioncode/library-managment-system/internal/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("Authorization")
		if tokenString == "" || err != nil {
			c.JSON(401, gin.H{"error message": "No Authorization header"})
			c.Abort()
			return
		}
		err, _ = utils.ValidateToken(tokenString)

		if err != nil {
			c.JSON(401, gin.H{"error message": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
