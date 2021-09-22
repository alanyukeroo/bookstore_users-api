package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/alanyukeroo/bookstore_users-api/domain/users"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))

	if err != nil {
		//TODO: Handle error
		return
	}
	fmt.Println(err)

	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO :Handle json error
		return
	}

	fmt.Println(user)

	c.String(http.StatusNotImplemented, "Ini ada POST Create User")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Ini adalah GET Search User")
}
