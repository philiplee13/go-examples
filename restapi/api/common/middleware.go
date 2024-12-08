package common

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CustomMiddleware(c *gin.Context) {
	fmt.Println("Inside custom middleware")

	// setting custom request id
	c.Writer.Header().Set("X-Request-Id", uuid.New().String())

	// validate token
	validateHeader(c)

	fmt.Println("sending request")

	c.Next()

	fmt.Println("after request")
}

// we could also have middleware that looks like this
// we could basically do some custom logic checking and then return
func CustomMiddleWareWithLogic() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		fmt.Println("after custom logic")
		c.Next()
	}
}

func validateHeader(c *gin.Context) {
	token := c.GetHeader("api-key")
	if token == "" {
		fmt.Println("token was not present")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Missing Api-Key"})
	}
}
