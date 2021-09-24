package users

import (
	"net/http"

	"github.com/alanyukeroo/bookstore_users-api/domain/users"
	"github.com/alanyukeroo/bookstore_users-api/services"
	"github.com/alanyukeroo/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func SearchUser(c *gin.Context) {

	c.String(http.StatusNotImplemented, "Ini adalah GET Search User")
}
