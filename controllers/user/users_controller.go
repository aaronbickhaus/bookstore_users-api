package user

import (
	"github.com/aaronbickhaus/bookstore_users-api/domain/users"
	"github.com/aaronbickhaus/bookstore_users-api/services"
	"github.com/aaronbickhaus/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Get(c *gin.Context) {
	userId, userErr := getUserId(c.Param("user_id"))

	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}

func Create(c *gin.Context) {
	 var user  users.User
	 if err := c.ShouldBindJSON(&user); err != nil {
  		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		 return
	 }

	 result, saveErr := services.UsersService.CreateUser(user)

	 if saveErr != nil {
	 	c.JSON(saveErr.Status, saveErr)
	 	return
	 }
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-PUBLIC") == "true"))

}

func Delete(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))

	if idErr != nil {
		idErr := errors.NewBadRequestError("invalid user id")
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	result, err := services.UsersService.SearchUsers(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}

func Update(c *gin.Context) {

	userId, idErr := getUserId(c.Param("user_id"))

	if idErr != nil {
		idErr := errors.NewBadRequestError("invalid user id")
		c.JSON(idErr.Status, idErr)
		return
	}

	var user  users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UsersService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
    c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-PUBLIC") == "true"))
}

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)

	if userErr != nil {
		return 0,  errors.NewBadRequestError("invalid user id")
	}
	return userId, nil
}
