package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user user.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(bytes))
	fmt.Println(err)
	fmt.Println(user)

	if err != nil {
		//TODO: Handle erro
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		// TODO :Handle jeson error
		return
	}
	c.String(http.StatusNotImplemented, "Ini ada POST Create User")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Ini adalah GET Search User")
}
