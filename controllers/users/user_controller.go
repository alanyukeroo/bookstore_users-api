package users

import (
	"fmt"
	"net/http"

	"github.com/alanyukeroo/bookstore_users-api/domain/users"
	"github.com/alanyukeroo/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		return
	}
	// Bisa digantikan ShouldBindJSON
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(bytes))

	// if err != nil {
	// 	//TODO: Handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	// TODO :Handle json error
	// 	return
	// }

	fmt.Println(user)

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: Handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Ini adalah GET Search User")
}
