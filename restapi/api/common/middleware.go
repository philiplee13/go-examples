package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CustomMiddleware(c *gin.Context) {
	fmt.Println("Inside custom middleware")

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
