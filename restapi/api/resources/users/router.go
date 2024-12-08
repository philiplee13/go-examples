package users

import "github.com/gin-gonic/gin"

type UserRouter struct {
	Controller UserController
	route      *gin.Engine
}

func (ur UserRouter) UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.GET("/", ur.Controller.GetAll)
		user.GET("/:id", ur.Controller.GetUserById)
		user.POST("/user", ur.Controller.Create)
		user.DELETE("/:id", ur.Controller.Delete)
		user.PUT("/:id", ur.Controller.Update)
	}
}
