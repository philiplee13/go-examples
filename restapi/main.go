package main

import (
	"context"
	"fmt"
	"restapi/api/resources/users"
	"restapi/config"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	dbconfig := config.Dbconfig{
		Host:     "localhost",
		Password: "postgres",
		User:     "postgres",
		Port:     5433,
	}
	conn, err := config.Connect(ctx, dbconfig)
	if err != nil {
		fmt.Println("error happened when trying to connect - exiting...")
	}

	router := gin.Default()

	uc := users.UserController{Database: conn}
	ur := users.UserRouter{Controller: uc}
	ur.UserRoutes(router)

	router.Run("localhost:8080")

	defer conn.Close(ctx)
}
