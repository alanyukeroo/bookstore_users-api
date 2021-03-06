package users

import (
	"fmt"

	"github.com/alanyukeroo/bookstore_users-api/utils/date_utils"
	"github.com/alanyukeroo/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func something() {
	user := User{}
	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.FirstName)
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprint("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprint("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprint("user %s already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
