package middleware

import (
	helper "golang-ecommerce-server/helpers"
	"golang-ecommerce-server/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Data: map[string]interface{}{"error": "UnAuthorized"}})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, responses.UserResponse{Status: http.StatusUnauthorized, Data: map[string]interface{}{"error": "UnAuthorized"}})
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Set("phone", claims.Phone)
		c.Set("email", claims.Email)
		c.Set("name", claims.Name)
		c.Next()
	}
}
