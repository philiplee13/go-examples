package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type UserController struct {
	Database *pgx.Conn
	us       UserService // instead of passing it in - instantitate on controller
}

func NewUserController(db *pgx.Conn) *UserController {
	us := UserService{}
	return &UserController{Database: db, us: us}
}

func (uc UserController) GetAll(c *gin.Context) {
	fmt.Println("inside get all method")
	users, err := uc.us.GetUsers(c, uc.Database)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": err})
	}
	c.JSON(http.StatusOK, users)
}

func (uc UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)

	user, err := uc.us.FindUser(c, uc.Database, userid)
	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": "Error happened when trying to find single user"})
	} else if user == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User not found"})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (uc UserController) Create(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("Error happened when trying to bind user %v\n", err)
		c.JSON(http.StatusExpectationFailed, gin.H{"message": err.Error()})
	}
	result := uc.us.CreateUser(c, uc.Database, user)
	if !result {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": "Error happened when trying to create user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User was created"})
}

func (uc UserController) Update(c *gin.Context) {
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)

	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Printf("Error happened when trying to bind user %v\n", err)
		c.JSON(http.StatusExpectationFailed, gin.H{"message": err.Error()})
	}
	result := uc.us.UpdateUser(c, uc.Database, userid, user)

	if !result {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": "Error happened when trying to update user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User was updated"})
}

func (uc UserController) Delete(c *gin.Context) {
	id := c.Param("id")
	userid, _ := strconv.Atoi(id)

	result := uc.us.DeleteUser(c, uc.Database, userid)

	if !result {
		c.JSON(http.StatusExpectationFailed, gin.H{"message": "Error happened when trying to delete user"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User was deleted"})
}
